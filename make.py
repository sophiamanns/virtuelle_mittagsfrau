#!/usr/bin/env python
# coding: utf-8

import os
from gluish.utils import shellout

filename_safe = "_-.0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

def resize_images(folder="data/DatenFlachs_Bilder",
                  size="28x28", output_directory="cache/images"):
    """
    Resize all images with `convert`.
    """
    if not os.path.exists(output_directory):
        os.makedirs(output_directory)

    for root, _, files in os.walk(folder):
        for i, fn in enumerate(sorted(files)):
            path = os.path.join(root, fn)

            output_filename = "size-%s-%03d.jpg" % (size, i)
            output_filename = "".join(c for c in output_filename if c in filename_safe)
            output_path = os.path.join(output_directory, output_filename)

            if os.path.exists(output_path):
                continue

            shellout("""convert "{input}" -resize {size} "{output}" """,
                     input=path, size=size, output=output_path)

if __name__ == '__main__':
    # resize_images(size="28x28")
    # resize_images(size="64x64\!")
    resize_images(size="256x256\!")
