# Ideas

* use some JS animation library
* try to build a small slot-gif thingy, that takes a bunch of images and creates gif animation with random outcomes

```
$ gifautomata *jpg > slotmachine.gif
```

Generate 250 images:

```
$ gifautomata -prefix sm -n 250 *jpg
$ ls
sm-000.gif sm-001.gif ... sm-250.gif
```

Then create a simple web service, where on each page load a new gif is
selected. People can write a short story in a text input and save it and
optionally publish or share it privately (including the image sequence).


# Slot machines

* http://odhyan.com/slot/

