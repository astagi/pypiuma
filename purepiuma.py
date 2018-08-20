import os
import time


from PIL import Image


def optimize(path, width, height):
    img = Image.open(path)
    img = img.resize((width, height), Image.ANTIALIAS)
    dest_path = os.path.join('.', 'piloptimized', os.path.basename(path))
    img.save(dest_path)
    img.close()
    return dest_path

images_to_optimize = [ item for item in os.listdir("images") for repetitions in range(100) ]

start = time.time()

for image_file_name in images_to_optimize:
    try:
        optimize(os.path.join("images", image_file_name), 100, 50)
    except Exception as err:
        print (err)

end = time.time()
print(end - start)
