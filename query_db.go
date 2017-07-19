package main

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/lib/pq" //http://godoc.org/github.com/lib/pq
    "log"
    "time"
    "path/filepath"
)

const (
    noiseCountSQL = `
SELECT
    loc.station,
    'pga-' || pga.vertical AS vertical,
    count(pga.*)
FROM
	impact.pga pga
	RIGHT OUTER JOIN impact.source loc ON loc.sourcepk = pga.sourcepk
GROUP BY
	loc.station, 'pga-' || pga.vertical
HAVING count(pga.*) > 16
UNION
SELECT
	loc.station,
    'pgv-' || pgv.vertical,
    count(pgv.*)
FROM
	impact.pgv pgv
	RIGHT OUTER JOIN impact.source loc ON loc.sourcepk = pgv.sourcepk
GROUP BY
	loc.station, 'pgv-' || pgv.vertical
ORDER BY 3 desc
        LIMIT 10`

    ratioDiffSQL = `
SELECT
        loc.station,
	CASE WHEN max_vert.max_pga > max_hori.max_pga THEN max_vert.max_pga / max_hori.max_pga ELSE max_hori.max_pga / max_vert.max_pga END ratio,
        max_vert.max_pga AS max_vertical,
        max_hori.max_pga AS max_horizontal
FROM
(
        SELECT
		sourcepk,
    		ROUND(MAX(pga), 8) AS max_pga
    	FROM
		impact.pga
       	WHERE
        	vertical = true
       	GROUP BY
        	sourcepk
) max_vert INNER JOIN
(
        SELECT
		sourcepk,
    		ROUND(MAX(pga), 8) AS max_pga
    	FROM
		impact.pga
       	WHERE
        	vertical = false
       	GROUP BY
        	sourcepk
) max_hori ON max_vert.sourcepk = max_hori.sourcepk
RIGHT OUTER JOIN impact.source loc ON loc.sourcepk = max_hori.sourcepk
ORDER BY
    	ratio DESC NULLS LAST
LIMIT 10`
)

var (
    trace *log.Logger
    db *sql.DB
    dir string
)

func init() {
    var err error
    file, err := os.OpenFile("/tmp/testdbquery.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Failed initializing logfile:", err)
    }

    trace = log.New(file, " ", log.LstdFlags|log.Lshortfile)

    t := time.Now()
    dir := filepath.Join("/tmp/",t.Format("20060102"))
    trace.Println("Files in: " + dir)
}

func main() {
        passwd, ok := os.LookupEnv("HAZARD_PASSWD")
        if !ok {
                trace.Fatalf("HAZARD_PASSWD not set for environment")
                os.Exit(1)
        }
        db, err := sql.Open("postgres",
                "postgres://hazard_r:" + passwd + "@geonet-api-ng-read.ccuclj9uvil4.ap-southeast-2.rds.amazonaws.com/hazard?sslmode=disable")

        if err != nil {
                trace.Fatalf("Error: problem with DB config: %s", err)
        }
        defer db.Close() // Pretty cool

        trace.Println("Getting top noise counts for Strong Motion")
        noiseCount(db)

        trace.Println("Getting ration for PGV")
        ratioDiff(db)
}

/* https://wiki.geonet.org.nz/display/dmcops/Strong+Motion+Noise+checks#StrongMotionNoisechecks-ConstantReportingCountNoise */
func noiseCount(db *sql.DB) {
    rows, err := db.Query(noiseCountSQL)

    if err != nil {
        trace.Fatalf("Error: %s", err)
    }

    var (
        station string
        component string
        count int
    )

    // There is almost a line for each station, open and closing for each line is about right

    for rows.Next() {
        err := rows.Scan(&station, &component, &count)
        if err != nil {
            trace.Fatalf("Error Scanning rows:%s", err)
        }
        // Just dump to log for now
        //trace.Printf("%s,%s,%d", station, component, count)

        //TODO: not working how I would expect
        dir := filepath.Join(dir, station)
        trace.Printf("Dir:", dir)
        if !(checkPath(dir)) {
                err := os.MkdirAll(dir, os.ModePerm)
                if err != nil {
                        trace.Printf("Failed creating dir:", err)
                }
        }
        file, err := os.OpenFile(filepath.Join(dir,"noiseCount.csv"), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
        if err != nil {
            trace.Printf("Failed opening file:", err)
        }
        defer file.Close()

        file.WriteString(fmt.Sprintf("%s,%d", component, count))
        /*if err != nil {
                panic(err)
        }*/

    }
}

/* https://wiki.geonet.org.nz/display/dmcops/Strong+Motion+Noise+checks#StrongMotionNoisechecks-PGAVerticalversusPGAHorizontalRatioNoise */
func ratioDiff(db *sql.DB) {
    rows, err := db.Query(ratioDiffSQL)

    if err != nil {
        trace.Fatalf("Error: %s", err)
    }

    var (
        station string
        ratio float64
        maxVertical float64
        maxHorizontal float64
    )

    trace.Printf("station,ratio,maxVertical,maxHorizontal")
    for rows.Next() {
            err := rows.Scan(&station, &ratio, &maxVertical, &maxHorizontal)
            if err != nil {
                   trace.Fatalf("Error Scanning rows:%s", err)
            }
            trace.Printf("%s,%f,%f,%f", station, ratio, maxVertical, maxHorizontal)
    }
}

func checkPath(path string) (exists bool) {
        _, err := os.Stat(path)

        if os.IsNotExist(err) {
                if err != nil {
                        return false
                }
        }
        return true
}
