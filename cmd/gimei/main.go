package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mattn/go-gimei"
)

func doName(name *gimei.Name, args []string) {
	for _, arg := range args {
		ret := ""
		switch arg {
		case "string":
			ret = name.String() // 斎藤 陽菜
		case "kanji":
			ret = name.Kanji() // 斎藤 陽菜
		case "hiragana":
			ret = name.Hiragana() // さいとう はるな
		case "katakana":
			ret = name.Katakana() // サイトウ ハルナ
		case "last-kanji":
			ret = name.Last.Kanji() // 斎藤
		case "last-hiragana":
			ret = name.Last.Hiragana() // さいとう
		case "last-katakana":
			ret = name.Last.Katakana() // サイトウ
		case "first-kanji":
			ret = name.First.Kanji() // 陽菜
		case "first-hiragana":
			ret = name.First.Hiragana() // はるな
		case "first-katakana":
			ret = name.First.Katakana() // ハルナ
		case "is-male":
			ret = fmt.Sprint(name.IsMale()) // false
		case "is-female":
			ret = fmt.Sprint(name.IsMale()) // false
		}
		fmt.Print(ret)
	}
}

func doAddress(address *gimei.Address, args []string) {
	for _, arg := range args {
		ret := ""
		switch arg {
		case "string":
			ret = address.String() // 岡山県大島郡大和村稲木町
		case "kanji":
			ret = address.Kanji() // 岡山県大島郡大和村稲木町
		case "hiragana":
			ret = address.Hiragana() // おかやまけんおおしまぐんやまとそんいなぎちょう
		case "katakana":
			ret = address.Katakana() // オカヤマケンオオシマグンヤマトソンイナギチョウ
		case "prefecture-string":
			ret = address.Prefecture.String() // 岡山県
		case "prefecture-kanji":
			ret = address.Prefecture.Kanji() // 岡山県
		case "prefecture-hiragana":
			ret = address.Prefecture.Hiragana() // おかやまけん
		case "prefecture-katakana":
			ret = address.Prefecture.Katakana() // オカヤマケン
		case "town-string":
			ret = address.Town.String() // 大島郡大和村
		case "town-kanji":
			ret = address.Town.Kanji() // 大島郡大和村
		case "town-hiragana":
			ret = address.Town.Hiragana() // おおしまぐんやまとそん
		case "town-katakana":
			ret = address.Town.Katakana() // オオシマグンヤマトソン
		case "city-string":
			ret = address.City.String() // 稲木町
		case "city-kanji":
			ret = address.City.Kanji() // 稲木町
		case "city-hiragana":
			ret = address.City.Hiragana() // いなぎちょう
		case "city-katakana":
			ret = address.City.Katakana() // イナギチョウ
		}
		fmt.Print(ret)
	}
}

func main() {
	var t string
	flag.StringVar(&t, "type", "name", "type of generator(name/male/female/address)")
	flag.Parse()

	switch t {
	case "name":
		doName(gimei.NewName(), flag.Args())
	case "male":
		doName(gimei.NewMale(), flag.Args())
	case "female":
		doName(gimei.NewFemale(), flag.Args())
	case "address":
		doAddress(gimei.NewAddress(), flag.Args())
	default:
		flag.Usage()
		os.Exit(2)
	}
	fmt.Println()
}
