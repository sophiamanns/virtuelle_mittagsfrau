# Die virtuelle Mittagsfrau

Für [Coding da Vinci Ost 2018](https://codingdavinci.de/events/ost/).

> Die [Mittagsfrau](https://de.wikipedia.org/wiki/Mittagsfrau) ist ein Naturgeist in weiblicher Gestalt in der slawischen Sagenwelt.

> Die pŕezpołnica hatte den serp (Sichel) in der Hand, und sagte, wenn jemand
> mittags auf dem Felde war: "Serp a šyju, Sichel und Hals". Und wer nicht eine
> Stunde lang erzählen konnte, dem hat sie den Kopf abgehauen"

![Lady Midday](images/Briefmarke_Sorbische_Sagen_-_Mittagsfrau_und_Nochtenerin_Crop.jpg)

## Daten

* https://speicherwolke.uni-leipzig.de/index.php/s/C99LbgXSSPdgFxV
* [Lokale Kopie](https://github.com/sophiamanns/virtuelle_mittagsfrau/tree/master/data)

![XXI 12_008](https://raw.githubusercontent.com/sophiamanns/virtuelle_mittagsfrau/master/data/DatenFlachs_Bilder/XXI%2012_008.jpg)

109 Bilder.


# TODO

* [ ] Excel Tags
* [ ] SLUB Fotothek Landscape, Geolocation, Cities (20-30)
* [ ] Webapp (Index, Details)
* [ ] Database
* [ ] SlotMachine-Details, responsive design ([v0](https://i.imgur.com/JFnEcS3.gif"))
* [ ] Hosting
* [ ] First translations (Anja, Wito)
* [ ] More stories
* [ ] Project texts and usage scenarios, for site and presentation

----

### Geografika der Bilddaten

* siehe [Metadate mit Tags](https://github.com/sophiamanns/virtuelle_mittagsfrau/blob/master/data/Metadaten_mit_Tags.csv)
* ad-hoc
  [Parser](https://github.com/sophiamanns/virtuelle_mittagsfrau/blob/master/fotothek.go)
für Daten der [Deutschen Fotothek](http://www.deutschefotothek.de/), mit
kleiner Auswahl an
[keywords](https://github.com/sophiamanns/virtuelle_mittagsfrau/blob/ad71fbfe4ba25ec705186b956852563031272e9c/fotothek.go#L155-L168),
mit
[3868](https://raw.githubusercontent.com/sophiamanns/virtuelle_mittagsfrau/master/data/fotothek.jsonl)
Matches.

```shell
$ csvcut -c7 Metadaten_mit_Tags.csv | sort | uniq -c | sort -nr
     36 Zabrod / Sabrodt
     16 Židźino / Seidewinkel
     13 wokoło Wojerecy / um Hoyerswerda
     10 Parcow / Groß Partwitz
      4 Sprjowje / Sprey
      4 Krakecy / Kreckwitz
      3 Rowno / Rohne
      3 Delnja Łužica / Niederlausitz
      3 Bluń / Bluno
      3
      2 Łusk pola Bukec / Lauske bei Hochkirch
      2 Lejno / Geierswalde
      2 Dešno / Dissen
      1 Wulke Zdźary / Groß Särchen
      1 Popojce / Papitz
      1 Ort sorbisch / Ort deutsch
      1 Nowe Město / Neustadt
      1 Janšojce / Jänschwalde
      1 Hausrat - Hausarbeit
      1 Groß Partwitz/ Parcow
      1 Bórkowy / Burg-Spreewald
      1 Błóta / Spreewald
```
