from PIL import Image, ImageOps
image = Image.open("../gen/image0.jpg")
size = (640, 640)
image.thumbnail(size, Image.ANTIALIAS)
background = Image.new('RGB', size, (255, 255, 255))
background.paste(
    image, (int((size[0] - image.size[0]) / 2),
            int((size[1] - image.size[1]) / 2))
)
background.save("output.jpg")
