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

* [x] [Excel Tags](https://github.com/sophiamanns/virtuelle_mittagsfrau/tree/master/data/Fotothek)
* [x] SLUB Fotothek Landscape, Geolocation, Cities (20-30), [data/Fotothek](https://github.com/sophiamanns/virtuelle_mittagsfrau/tree/master/data/Fotothek)
* [x] Webapp (Index, Details)
* [x] Database
* [x] SlotMachine-Details, responsive design ([v0](https://i.imgur.com/JFnEcS3.gif"), [v1](https://i.imgur.com/hm5CSz8.mp4))
* [x] Hosting, [@scaleway](http://51.15.235.18), [mittagsfrau.de](http://mittagsfrau.de)
* [x] First translations (Anja, Wito)
* [ ] More stories
* [x] Project texts and usage scenarios, for site and presentation

----

### Geografika der Bilddaten

* siehe [Metadate mit Tags](https://github.com/sophiamanns/virtuelle_mittagsfrau/blob/master/data/Metadaten_mit_Tags.csv)
* ad-hoc
  [Parser](https://github.com/sophiamanns/virtuelle_mittagsfrau/blob/master/fotothek.go)
für Daten der [Deutschen Fotothek](http://www.deutschefotothek.de/), mit
kleiner Auswahl an
[keywords](https://github.com/sophiamanns/virtuelle_mittagsfrau/blob/3a35b323078bdb62dccb0e42383527f113e9b78f/fotothek.go#L22-L36)
mit
[3868](https://raw.githubusercontent.com/sophiamanns/virtuelle_mittagsfrau/master/data/fotothek.jsonl)
matches.

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

----

Bildgruppen: Menschen, Artefakte, Landschaften.

```python
In [8]: df = pd.read_csv("https://git.io/vhrG8")

In [9]: df["Tag"].value_counts()
Out[9]:
Mensch        49
Technik       48
Landschaft    12
Name: Tag, dtype: int64
```

Weitere 38 Bilder aus der Fotothek für Gruppe Landschaft (draußen).

Verteilung der Bilder:

* Mensch: 49
* Technik: 48
* Landschaft: 50

Theoretische Anzahl der Kombinationen: 117600. Einschränkung auf 15^3
Beispiele, da wir die Animationen vorberechnen - und jedes mp4 ca. 2-4M
benötigt ~ 3375 * 4 ~ 13G.

