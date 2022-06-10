import sys
import ppa6
import PIL.Image as Image
import numpy as np

if (len(sys.argv) < 1):
    print("No argument given")
    sys.exit(1)

printer = ppa6.Printer('35:53:19:07:1D:BC', ppa6.PrinterType.A6p)
printer.connect()
printer.reset()

print(f'Name: {printer.getDeviceName()}')
print(f'S/N: {printer.getDeviceSerialNumber()}')
print(f'F/W: {printer.getDeviceFirmware()}')
print(f'Battery: {printer.getDeviceBattery()}%')
print(f'H/W: {printer.getDeviceHardware()}')
print(f'MAC: {printer.getDeviceMAC()}')
print(f'Full: {printer.getDeviceFull()}')

printer.setConcentration(1)

c_array = np.asarray(Image.open(sys.argv[1]).convert('L'))
im_1 = Image.fromarray(c_array)
image = im_1
size = (640, 640)
image.thumbnail(size, Image.ANTIALIAS)
background = Image.new('RGB', size, (255, 255, 255))
background.paste(
    image, (int((size[0] - image.size[0]) / 2),
            int((size[1] - image.size[1]) / 2))
)
im_1 = background

c_array = np.asarray(Image.open(sys.argv[2]).convert('L'))
im_2 = Image.fromarray(c_array)
image = im_2
size = (640, 640)
image.thumbnail(size, Image.ANTIALIAS)
background = Image.new('RGB', size, (255, 255, 255))
background.paste(
    image, (int((size[0] - image.size[0]) / 2),
            int((size[1] - image.size[1]) / 2))
)
im_2 = background

c_array = np.asarray(Image.open(sys.argv[3]).convert('L'))
im_3 = Image.fromarray(c_array)
image = im_3
size = (640, 640)
image.thumbnail(size, Image.ANTIALIAS)
background = Image.new('RGB', size, (255, 255, 255))
background.paste(
    image, (int((size[0] - image.size[0]) / 2),
            int((size[1] - image.size[1]) / 2))
)
im_3 = background

c_array = np.asarray(Image.open(sys.argv[4]).convert('L'))
im_4 = Image.fromarray(c_array)
image = im_4
size = (640, 640)
image.thumbnail(size, Image.ANTIALIAS)
background = Image.new('RGB', size, (255, 255, 255))
background.paste(
    image, (int((size[0] - image.size[0]) / 2),
            int((size[1] - image.size[1]) / 2))
)
im_4 = background

f = open("/phoenix_images/scene.txt", "r")
scene = f.read()
seed = scene.split("{")[1]

printer.writeASCII('Seed ' + seed + '\n')
printer.printBreak(10)
printer.printImage(im_1)
printer.printBreak(10)
printer.printImage(im_2)
printer.printBreak(10)
printer.printImage(im_3)
printer.printBreak(10)
printer.printImage(im_4)
printer.printBreak(100)
f.close()
