
def main():
    b = [0]
    c = 0
    for i in b:
        b.append(i+1)
        c += 1
        if c == 100: break
    print("===> %s"%b)

    a = {0: 0}
    for i in a:
        a[i+1] = i+1

if __name__ == "__main__":
    main()