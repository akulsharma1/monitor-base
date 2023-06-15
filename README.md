# polymarket-monitor-public

**Functionality**

This repository allows you to monitor the [Polymarket](https://polymarket.com/) website for new bets.
The code also has modular infrastructure in place to add new websites to monitor extremely quickly.

**Importance of the monitor**

Polymarket's betting system started off at 50/50 odds for each side, and the more people that invested, the more the odds changed.
This meant that if you were to bet at 50/50 odds, and the odds eventually became 90/10 in your favor, you could sell your position for a profit before the bet completed, thus resulting in no risk.
This monitor allowed us to be the first to snipe odds, allowing one to consistently snipe odds at 50/50 and profit without risk.

**How to run the monitor**

1. Clone the repository
2. Add your user in [users.json](users.json)
3. run `go run .`
4. The monitor will then run and send new polymarket listings to the webhook you provided in `users.json`.


**Why Open Source?**

Recently, the Polymarket website has had far fewer bets available, and the amount of liquidity in each bet has been lower. Very few bets reach an amount where it is consistently profitable to run the monitor and bet.
Now, I have opensourced the code for beginners to learn from.

**Future Plans**

Integrate new websites into the monitor (e.g. my [shopify store monitor](https://github.com/akulsharma1/publicshopifymonitor))