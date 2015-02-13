package other

import (
	"github.com/advancedlogic/GoOse"
)

func main() {
	g := goose.New()
	article := g.ExtractFromUrl("http://edition.cnn.com/2012/07/08/opinion/banzi-ted-open-source/index.html")
	println("title", article.Title)
	println("description", article.MetaDescription)
	println("keywords", article.MetaKeywords)
	println("content", article.CleanedText)
	println("url", article.FinalUrl)
	println("top image", article.TopImage)
}
