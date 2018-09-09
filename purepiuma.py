import os
import subprocess
import time

from threading import Thread

from PIL import Image


FNULL = open(os.devnull, 'w')


def convert(img_format, dest_path):
    if img_format == 'JPEG':
        opt_command = 'jpegoptim "{0}" --max=100'.format(dest_path)
    elif img_format == 'PNG':
        opt_command = 'pngquant "{0}" -f --ext=.png'.format(dest_path)
    else:
        raise Exception('No file image supported')
    subprocess.Popen(
        [opt_command], shell=True,
        stdout=FNULL, stderr=subprocess.STDOUT
    ).wait()


def optimize(path, width=0, height=0):
    try:

        dest_path = os.path.join('.', 'piloptimized', os.path.basename(path))
        img = Image.open(path)
        img_format = img.format

        if height == 0:
            height = int(width * img.size[1] / img.size[0])
        elif width == 0:
            width = int(height * img.size[0] / img.size[1])

        img = img.resize((width, height), Image.ANTIALIAS | Image.NEAREST)
        img.save(dest_path)
        img.close()
        convert(img_format, dest_path)
        return dest_path
    except Exception as err:
        print (err)


def pureppiuma():
    images_to_optimize = [ item for item in os.listdir("images") for repetitions in range(1) ]
    for image_file_name in images_to_optimize:
        optimize(os.path.join("images", image_file_name), 200)


start = time.time()
pureppiuma()
end = time.time()
print(end - start)
