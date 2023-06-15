# monitor-base

**Functionality**

This repository contains the base infrastructure and an example website [Polymarket](https://polymarket.com/) which it monitors.
You can quickly add new websites using the base infrastructure as well.
I initially used this code to monitor the polymarket website, and later expanded it to become a more general codebase for public use.

**Importance of the monitor**

Polymarket's betting system started off at 50/50 odds for each side, and the more people invested, the more the odds changed.
This meant that if you were to bet at 50/50 odds, and the odds eventually became 90/10 in your favor, you could sell your position for a profit before the bet was completed, thus resulting in no risk.
This monitor allowed us to be the first to snipe odds, allowing one to snipe odds at 50/50 and profit without risk consistently.

**How to run the monitor**

1. Clone the repository
2. Add your user in [users.json](users.json)
3. run `go run .`
4. The monitor will then run and send new polymarket listings to the webhook you provided in `users.json`.


**Why Open Source?**

Recently, the Polymarket website has had far fewer bets available, and the amount of liquidity in each bet has been lower. Very few bets reach an amount where it is consistently profitable to run the monitor and bet.
Now, I have open-sourced the code for beginners to learn from.

**Future Plans**

Integrate new websites into the monitor (e.g. my [shopify store monitor](https://github.com/akulsharma1/publicshopifymonitor))
