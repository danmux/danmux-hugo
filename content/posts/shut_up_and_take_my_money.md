---
title: "Shut Up and Take My Money Summary"
description: "A summary of the key points from Vincent Hauperts talk at the Chaos Communication Congress."
date: "2016-12-30"
categories: 
    - "security"
    - "tech"
    - "fintech"
    - "engineering"
---

[Vincent Haupert](https://fahrplan.events.ccc.de/congress/2016/Fahrplan/speakers/5995.html) Presented at the Chaos Communication Congress A talk about some severe security breaches in the fintech startup [Number26](https://n26.com/) app and API. Here is a summary of the key points.

<!--more-->

If you only have 5 mins to grab the key points they are summarised below. 

The Original presentation is [here](https://media.ccc.de/v/33c3-7969-shut_up_and_take_my_money) and it is quite fun to watch. 

Or watch the Youtube version here:

<iframe width="560" height="315" src="https://www.youtube.com/embed/KopWe2ZpVQI" frameborder="0" allowfullscreen></iframe>


Summary
=======

1. **No certificate pinning**, making MitM of the API, much easier.
2. **Password reset via email only** - compromised email account may be more likely if no 2fa and easy password (but actually access to the email account is not even needed - *see later!* )
3. **Spear fishing is made more easy** than it should be to gain login credentials.
	* 3.1. API accepts (un-hashed) email and responds if they are a Numbr26 user or not (to support address book based send money).
	* 3.2. The email lookup API is not throttled.
	* 3.3. 68M dropbox emails evaluated against the API, no blocking, no suspicion raised at all.
	* 3.4. 33k users identified - nice list for targeted (spear) phishing.
4. **Weak pairing scheme** the token sharing, phone pairing scheme is *only a client side feature* (no keys checked on server), so the raw API call for send money works from any client!
5. **Send money not rate limited.** Sent 2000 1c transactions to the same user in 30 mins, no suspicion raised. 
    * 5.a. 3 weeks later N26 contacted the *recipient* (not the sender) and wanted to block his account! 
6. **Email reset token leaked in API.** The reset API request responds with the actual token/link (should only be sent to the email account), so you don't actually need the email account to reset credentials!
7. **Secret ID Leaked in the API.** Other secure features need a 'secret' card ID that is readily discoverable. This ID is not the 16 digit card number, but another ID/Token printed on the physical MasterCard. However this ID is also used in transaction ID's and is leaked in the API, hence possession of the card is not needed.
8. **Pin vulnerable to brute force**.
    * 8.1. The 5 digit Pin is needed for some secure requests.
    * 8.2. Pin check API is not throttled, managed 160 rps, so averagely takes 5 mins to brute force the PIN.
9. **Full account control.** With those 4 facts - total control of all account features is possible - e.g. unpairing a device.
11. **Unparing token leaked in API.** Un-pairing link also sent in API response (no email account needed)
12. **Support call authentication secret leaked in API.** Changing email of an account needs a support call (support can also unpair a device). Support need the Master Card ID, Account balance, Place Of Birth.
    * 12.1. Place Of birth, is also returned in an personal info API response, now attacker has all three facts and can change email, making further exploitation stealthier.
14. **Risk elevation.** Even accounts with no balance can be exploited because the attacker can get an overdraft of up to €2000 approved in 2 mins.

Comment
=======

These vulnerabilities are incredibly naive, and N26 has let the whole challenger FinTech community down. We should be building a reputation and trust with consumers, by blowing the banks out of the water, on all angles. 

To be this negligent with security is beyond comprehension...
----------------------------------------------------------

1. No rate limiting.
2. Leaking security facts in the API.
3. Leaking security tokens in the API.
4. No token validation on the API (client only)
5. No effective 2fa.
6. No hashing of emails in unsecured API.
7. No other unusual behavior detection.

### N.B.
Regarding N26 targeting the the recipient account for blocking, Haupert draws an analogy of blocking your email account for receiving spam. The difference here is that value is received. Perhaps N26 were wise to target the recipient as possibly being the more likely attacker. However it *was* the senders account that was compromised, and could still be used to transfer money into any number of other accounts. Both accounts should be secured. 

