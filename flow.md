// go to this
https://cskh.npc.com.vn/home/AccountNPC

// Login
await fetch("https://cskh.npc.com.vn/Account/Login", {
    "credentials": "include",
    "headers": {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:145.0) Gecko/20100101 Firefox/145.0",
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
        "Accept-Language": "en-US,en;q=0.5",
        "Content-Type": "application/x-www-form-urlencoded",
        "Upgrade-Insecure-Requests": "1",
        "Sec-Fetch-Dest": "document",
        "Sec-Fetch-Mode": "navigate",
        "Sec-Fetch-Site": "same-origin",
        "Sec-Fetch-User": "?1",
        "Idempotency-Key": "\"3124866401522697894\"",
        "Priority": "u=0, i"
    },
    "referrer": "https://cskh.npc.com.vn/home/AccountNPC",
    "body": "__RequestVerificationToken=tDGDW-3952MjwpsrgeHge1Ycx6H8XGvRHUKVOX_yLwT2hwCECE3siDL-aFENLrLgqLMsoI50YOvGiAy09535NPc7fX0ZR8nU9BRsRsoAKeA1&Username=PA04GT7017040&Password=123456&CaptchaDeText=a6b6a6d3ef3d42f3b1de9d017e538725&CaptchaInputText=NNGI&previousLink=",
    "method": "POST",
    "mode": "cors"
});

// Search
await fetch("https://cskh.npc.com.vn/HoaDon/TraCuuHDSPC?ky=1&thang=9&nam=2025&_=1761317429094", {
    "credentials": "include",
    "headers": {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:145.0) Gecko/20100101 Firefox/145.0",
        "Accept": "text/html, */*; q=0.01",
        "Accept-Language": "en-US,en;q=0.5",
        "Content-Type": "application/json; charset=utf-8",
        "X-Requested-With": "XMLHttpRequest",
        "Sec-Fetch-Dest": "empty",
        "Sec-Fetch-Mode": "cors",
        "Sec-Fetch-Site": "same-origin",
        "Priority": "u=0"
    },
    "referrer": "https://cskh.npc.com.vn/DichVuTTCSKH/IndexNPC?index=1",
    "method": "GET",
    "mode": "cors"
});
