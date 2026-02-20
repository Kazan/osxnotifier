package main

import "fyne.io/fyne/v2"

var bellTemplateIconSVG = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 22 22" fill="none">
  <path d="M11 4.2C8.84 4.2 7.1 5.95 7.1 8.1V10.5C7.1 11.89 6.53 13.2 5.54 14.15L4.9 14.76C4.72 14.93 4.67 15.2 4.78 15.42C4.89 15.64 5.11 15.78 5.36 15.78H16.64C16.89 15.78 17.11 15.64 17.22 15.42C17.33 15.2 17.28 14.93 17.1 14.76L16.46 14.15C15.47 13.2 14.9 11.89 14.9 10.5V8.1C14.9 5.95 13.16 4.2 11 4.2Z" fill="currentColor"/>
  <path d="M9.2 17.1C9.36 18.09 10.11 18.8 11 18.8C11.89 18.8 12.64 18.09 12.8 17.1H9.2Z" fill="currentColor"/>
</svg>
`)

func getIconResource() fyne.Resource {
	return fyne.NewStaticResource("bellTemplate.svg", bellTemplateIconSVG)
}
