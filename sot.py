#!/usr/bin/python3
import sys

tmp = []
tmp2 = []

def main():
    try:
        file = sys.argv[1]
        arr = open(file, 'r').read().split('\n')

        #adding the index of each first diffrent url to an array
        for i in range(len(arr)):
            string = arr[i].partition("://")[2]
            string = string.partition("/")[0]
            if string not in tmp2:
                tmp.append(i)
                tmp2.append(string)

        #doing the main job of this script
        for i in range(len(arr)):
            for j in tmp:
                if i + j >= len(arr):
                    break
                else:
                    current = arr[j].partition("://")[2].partition("/")[0]
                    next = arr[j + i].partition("://")[2].partition("/")[0]
                    c = arr[j]
                    n = arr[j + i]
                    if current == next:
                        print(arr[i + j])

        
    except Exception as e:
        print("Usage: sot [file path]")

main()
