import os
import threading
import time


from PIL import Image


def convert(img_format, dest_path):
    if img_format == 'JPEG':
        opt_command = 'jpegoptim {0}'.format(dest_path)
    elif img_format == 'PNG':
        opt_command = 'pngquant {0} -f --ext=.png'.format(dest_path)
    else:
        raise Exception('No file image supported')
    os.system(opt_command)

def optimize(path, width, height):
    img = Image.open(path)
    img_format = img.format
    img = img.resize((width, height), Image.ANTIALIAS)
    dest_path = os.path.join('.', 'piloptimized', os.path.basename(path))
    img.save(dest_path)
    img.close()
    thread = threading.Thread(target = convert, args = (img_format, dest_path) )
    thread.start()
    return dest_path

images_to_optimize = [ item for item in os.listdir("images") for repetitions in range(1) ]

start = time.time()

for image_file_name in images_to_optimize:
    try:
        optimize(os.path.join("images", image_file_name), 100, 50)
    except Exception as err:
        print (err)

end = time.time()
print(end - start)
