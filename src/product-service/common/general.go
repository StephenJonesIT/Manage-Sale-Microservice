package common

import (
        "fmt"
        "math/rand"
        "strconv"
        "strings"
        "time"
)

func GeneralProductCode(brand string) string {
        result := getFirstLetters(brand)
        result += getCurrentTime()
        result += generateRandomNumberString(3)
        return strings.ToUpper(result)
}

func getFirstLetters(s string) string {
        words := strings.Fields(s)
        result := ""
        for _, word := range words {
                if len(word) > 0 {
                        result += string(word[0])
                }
        }
        return result
}

func getCurrentTime() string {
        now := time.Now()
        year := now.Year()
        yearStr := strconv.Itoa(year)

        // Lấy ký tự đầu và cuối
        firstChar := yearStr[0]
        lastChar := yearStr[len(yearStr)-1]
        month := int(now.Month())
        day := now.Day()

        return fmt.Sprintf("%d%d%02d%02d", firstChar,lastChar, month, day)
}

func generateRandomNumberString(length int) string {
        rand.Seed(time.Now().UnixNano())
        var sb strings.Builder
        for i := 0; i < length; i++ {
                randomNumber := rand.Intn(10)
                sb.WriteString(strconv.Itoa(randomNumber))
        }
        return sb.String()
}