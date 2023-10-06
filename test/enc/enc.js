// https://gist.github.com/nikolahellatrigger/a8856463170fbe3596569977148ebaf4

/*
    This script is used to "encrypt" somes data into fingerprint_event such as:
        - webgl vendor + renderer
        - browser performance
        - browser timezone
    
    I think it's used to verify the data is authentic / non duplicated (output is different each time you run the function)
*/

function getRandNum(A, g) {
    return Math.floor(Math.random() * (g - A + 1)) + A
}

function __Enc(A) {
    if (null == A) {
        return null
    }

    let inputArr = Array.from({ length: 13 }, function () {
        return String.fromCharCode(getRandNum(65, 90))
    }).join('')

    let rand_a = getRandNum(1, 26)

    let encodedResult = A.split(' ').reverse().join(' ').split('').reverse().map(function (A) {
        if (!A.match('/[a-z]/i')) {
            return A
        }

        let I = hA.indexOf(A.toLowerCase())
        let B = hA[(I + rand_a) % 26]

        if (A === A.toUpperCase()) {
            return B.toUpperCase();
        } else {
            return B;
        }
    }).join('')

    let b64out = window.btoa(encodeURIComponent(encodedResult)).split('').reverse().join('')
    let b64randLen = getRandNum(1, b64out.length - 1)

    return [
        (b64out.slice(b64randLen, b64out.length) + b64out.slice(0, b64randLen)).replace(
            new RegExp('['.concat(inputArr).concat(inputArr.toLowerCase(), ']'), 'g'),
            
            function (A) {
                if (A === A.toUpperCase()) {
                    return A.toLowerCase();
                } else {
                    return A.toUpperCase();
                }
            }
        ),
        rand_a.toString(16),
        b64randLen.toString(16),
        inputArr,
    ]
}

__Enc("ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)")

// output
/*
[
    "wzWItjozVSElUqWItjO4kVjRUSbNkmlAJmlEktHxUR=q0mEftmPAJmlA3CFVZXWmkmlAJmly3CFVZXWAJmlqUaYV2y0NDRxEDmYUcVPbJmlmDm2ADmYUiUUhfmYUYRlz0BYN",
    "13",
    "5a",
    "FCOZWPYTZJMBQ"
]
*/