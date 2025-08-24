package main

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

func WriteSegmentsCSV(w io.Writer, segments []*Segment) error {
	wtr := csv.NewWriter(w)
	defer wtr.Flush()
	wtr.Write([]string{"key","description","percent","start_date","end_date"})
	for _, s := range segments {
		start, end := "", ""
		if s.StartDate != nil { start = s.StartDate.Format(time.RFC3339) }
		if s.EndDate != nil { end = s.EndDate.Format(time.RFC3339) }
		wtr.Write([]string{s.Key, s.Description, strconv.Itoa(s.Percent), start, end})
	}
	return nil
}

func ReadSegmentsCSV(r io.Reader) ([]*Segment, error) {
	rdr := csv.NewReader(r)
	records, err := rdr.ReadAll()
	if err != nil { return nil, err }
	var segs []*Segment
	for i, rec := range records {
		if i==0 { continue }
		percent,_ := strconv.Atoi(rec[2])
		var start, end *time.Time
		if rec[3] != "" { t,_ := time.Parse(time.RFC3339, rec[3]); start = &t }
		if rec[4] != "" { t,_ := time.Parse(time.RFC3339, rec[4]); end = &t }
		segs = append(segs, &Segment{
			Key: rec[0], Description: rec[1], Percent: percent,
			StartDate: start, EndDate: end,
			CreatedAt: time.Now(), UpdatedAt: time.Now(),
		})
	}
	return segs, nil
}
