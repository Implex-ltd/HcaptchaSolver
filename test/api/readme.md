# Task config
```js
{
    "domain": "example.com",
    "site_key": "0000000-0000-0000-0000-00000000000",
    "user_agent": "xx", // (@enterprise -> the ua that you are using to submit token)
    "proxy": "",        // http://user:pass@ip:port | http://ip:port (@enterprise -> the proxy must match the ip that submit captcha)
    "task_type": 0,     // 0= ENTERPRISE, 1= NORMAL
    "invisible": false, // is the captcha invisible ?
    "rqdata": rqdata,   // @enterprise -> rqdata payload 
    "a11y_tfe": false,  // use text captcha (must be enabled on the website) (@enterprise -> quality is lower)
    "turbo": true,      // enable turbo mode (submit task faster)
    "turbo_st": 3000,   // minimum time to submit captcha in ms
}
```

# Sitekey list
```
www.habbo.fr
edc4ce89-8903-4906-80b1-7440ad9a69c8

accounts.autodesk.com
636943a1-4920-4970-a0ad-42d4aff214ce

dashboard.stripe.com
89378a0b-0942-4717-89fc-52e01acddedd

www.hostinger.com
bd07a95b-c4b5-4bfc-98ed-c310c4df2370

gate.com.ph
03080def-874d-4bef-90e3-2f71c2c69202

comspec.com.ph
3d4e78fa-92a0-4b4b-b404-c76e112c4d02

sorial.pe
108d9b11-ddc2-4f49-9622-fb7c90144817

www.yourlifespeaks.org
13547e83-ad0b-4b77-ba7d-2f650809b31f

worldpittsburgh.org
2578257c-7771-4398-86c5-5f9d9571a2b2

wingardhome.org
222dfa5e-93ed-4ab2-a48f-c35eec04f2ad

tousauxabris.org
f362115e-54cb-4aa5-8bee-e964d9b71fdf

www.bitstamp.net
55358dd0-6380-4e69-8390-647a403a8a7f

www.herblaysurseine.fr
95d223ff-5af0-448a-9b16-567876393610
```
