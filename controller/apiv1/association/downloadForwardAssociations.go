package association

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"singlishwords/log"
	"singlishwords/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


func DownloadForwardAssociations(c *gin.Context) {
	word := c.Param("word")
	word = strings.Replace(word, "-", " ", -1)
	log.Logger.Infof(fmt.Sprintf("Getting forward associations for word: %s", word))

	_, associations, err := service.GetSetAndForwardAssociations(word)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

    buffer := &bytes.Buffer{} 
    writer := csv.NewWriter(buffer)
	writer.Write([]string{"source", "target", "count"})
    for _, as := range associations {
        err := writer.Write([]string{as.Source, as.Target, strconv.FormatInt(as.Count, 10)})
        if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
    }
    writer.Flush()

	contentLength := buffer.Len()
    contentType := "text/csv"
    extraHeaders := map[string]string{
      "Content-Disposition": `attachment;filename=associations.csv"`,
    }

    c.DataFromReader(http.StatusOK, int64(contentLength), contentType, buffer, extraHeaders)
}
