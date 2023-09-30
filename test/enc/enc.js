function __EncryptStr(A) {
    // console.log(A)
    var g = 322,
        I = 617,
        B = 619,
        Q = 378,
        C = 671,
        E = 671,
        D = 450,
        i = 768,
        w = 588,
        o = 242,
        M = h;
    if (null == A) return null;
    var N = {};
    N[M(238)] = 13;
    var G = M(352) != typeof A ? String(A) : A,
        y = Array[M(g)](N, (function () {
            return String.fromCharCode(cA(65, 90))
        }))[M(378)](""),
        a = cA(1, 26),
        n = G[M(619)](" ")[M(I)]()[M(378)](" ")[M(B)]("").reverse().map((function (A) {
            var g = M;
            if (!A[g(659)](YA)) return A;
            var I = hA[g(w)](A.toLowerCase()),
                B = hA[(I + a) % 26];
            return A === A[g(o)]() ? B[g(o)]() : B
        }))[M(Q)](""),
        L = window.btoa(encodeURIComponent(n)).split("")[M(I)]()[M(378)](""),
        c = L[M(238)],
        Y = cA(1, c - 1);
    return [(L[M(C)](Y, c) + L[M(E)](0, Y))[M(724)](new RegExp("["[M(450)](y)[M(D)](y[M(576)](), "]"), "g"), (function (A) {
        var g = M;
        return A === A[g(242)]() ? A[g(576)]() : A[g(242)]()
    })), a[M(768)](16), Y[M(i)](16), y]
}

__EncryptStr("ANGLE (NVIDIA, NVIDIA GeForce GT 755M Direct3D11 vs_5_0 ps_5_0, D3D11)")

// deob, (not 100%)

function __EncryptStr(input) {
    if (input === null) return null;

    const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    const randomOffset = Math.floor(Math.random() * 26);

    const encryptedString = input
        .split(' ')
        .map(word => {
            return word.split('').reverse().map(char => {
                if (!alphabet.includes(char.toUpperCase())) return char;
                const charCode = char.toLowerCase().charCodeAt(0);
                const encryptedCharCode = (charCode - 97 + randomOffset) % 26 + 97;
                return char === char.toUpperCase()
                    ? String.fromCharCode(encryptedCharCode).toUpperCase()
                    : String.fromCharCode(encryptedCharCode);
            }).join('');
        })
        .join(' ');

    const encodedString = btoa(encodeURIComponent(encryptedString))
        .split('')
        .map(char => {
            const isUppercase = /[A-Z]/.test(char);
            return isUppercase ? char.toLowerCase() : char.toUpperCase();
        })
        .join('');

    const length = encodedString.length;
    const randomIndex = Math.floor(Math.random() * (length - 1)) + 1;
    const firstPart = encodedString.substring(randomIndex);
    const secondPart = encodedString.substring(0, randomIndex);

    const finalString = (firstPart + secondPart)
        .replace(new RegExp(`[${alphabet}]`, 'gi'), char => {
            return char === char.toUpperCase() ? char.toLowerCase() : char.toUpperCase();
        });

    return [finalString, randomOffset.toString(16), randomIndex.toString(16)];
}