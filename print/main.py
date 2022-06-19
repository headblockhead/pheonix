import sys
import ppa6
import PIL.Image as Image
from subprocess import Popen, PIPE
import numpy as np
import qrcode
import requests
import json
from io import BytesIO
import base64
from time import sleep
import os


def get_printing_data(url):
    encoded_images_request = requests.get(url)
    encoded_images_request_json = json.loads(encoded_images_request.text)
    encoded_images = encoded_images_request_json['Frames']
    decoded_image_1 = BytesIO(base64.b64decode(encoded_images[0]))
    decoded_image_2 = BytesIO(base64.b64decode(encoded_images[1]))
    decoded_image_3 = BytesIO(base64.b64decode(encoded_images[2]))
    decoded_image_4 = BytesIO(base64.b64decode(encoded_images[3]))

    c_array = np.asarray(Image.open(decoded_image_1).convert('L'))
    im_1 = Image.fromarray(c_array)
    c_array = np.asarray(Image.open(decoded_image_2).convert('L'))
    im_2 = Image.fromarray(c_array)
    c_array = np.asarray(Image.open(decoded_image_3).convert('L'))
    im_3 = Image.fromarray(c_array)
    c_array = np.asarray(Image.open(decoded_image_4).convert('L'))
    im_4 = Image.fromarray(c_array)

    seed = encoded_images_request_json["Seed"]

    img = qrcode.make(
        "https://www.headblockhead.com/projects/phoenix/?seed=" + str(seed))

    image = img
    size = (120, 120)
    image.thumbnail(size, Image.ANTIALIAS)
    background = Image.new('RGB', size, (255, 255, 255))
    background.paste(
        image, (int((size[0] - image.size[0]) / 2),
                int((size[1] - image.size[1]) / 2))
    )
    img = background

    return ['Here is your randomly generated Phoenix Wright comic!\n', 10, 'Your Seed: ' + str(seed) + '\n', 20, im_1, 10, im_2, 10, im_3, 10, im_4, 30, 'To see more, scan this QR code: \n', img, 100]


if (len(sys.argv) < 1):
    print("No argument given")
    sys.exit(1)

print_mode = os.getenv('PHOENIX_PRINT_MODE')
if print_mode == 'decode':
    print("Mode: Decoding")
    url = sys.argv[1]
    data = get_printing_data(url)
    f = open('/phoenix_output/log.txt', 'w')
    log = []
    for i in range(len(data)):
        if i % 2 == 0:
            sleep(2)
        if (type(data[i]) == int):
            log.append("Break: " + str(data[i]) + '\n')
            print("Break: " + str(data[i]) + '\n')
        elif (type(data[i]) == str):
            log.append("ASCII: " + data[i] + '\n')
            print("ASCII: " + data[i])
        elif (type(data[i]) == Image.Image):
            filename = "/phoenix_output/PrintedImage_" + str(i) + ".png"
            data[i].save(filename)
            log.append("Image - located in " + filename + '\n')
            print("Image - located in " + filename + '\n')
    f.writelines(log)
    sys.exit(0)
elif print_mode == 'print':
    print("Mode: Printing")
    printer = ppa6.Printer('35:53:19:07:1D:BC', ppa6.PrinterType.A6)
    printer.connect()
    printer.reset()
    print(f'Name: {printer.getDeviceName()}')
    # print(f'S/N: {printer.getDeviceSerialNumber()}')
    # print(f'F/W: {printer.getDeviceFirmware()}')
    print(f'Battery: {printer.getDeviceBattery()}%')
    # print(f'H/W: {printer.getDeviceHardware()}')
    # print(f'MAC: {printer.getDeviceMAC()}')
    # print(f'Full: {printer.getDeviceFull()}')
    printer.setConcentration(1)
    url = sys.argv[1]
    data = get_printing_data(url)
    for i in range(len(data)):
        if i % 2 == 0:
            sleep(1.5)
        if (type(data[i]) == int):
            printer.printBreak(data[i])
        elif (type(data[i]) == str):
            printer.writeASCII(data[i])
        elif (type(data[i]) == Image.Image):
            printer.printImage(data[i])
    sys.exit(0)
