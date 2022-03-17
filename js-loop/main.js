

function main() {
    a = [0]
    c = 0
    for (var i = 0; i < a.length; i++) {
        a.push(i+1)
        c++
        if (c > 100) {
            break
        }
    }
    console.log("===>", a)
    b = {0: 0}
    for (const [key, value] of Object.entries(b)) {
        b[key+1] = key+1
      }
    console.log("===>", b)
}

main();