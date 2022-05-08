## Inspiration
Google assistant, Alexa and Siri are limited. I use them  mostly as a glorified kitchen timer.

They are limited as they are sandboxed, isolated because of privacy concerns.
Some of us keep journals, have personal assistants or have lawyers act on our behalf.
What if we could have a device or a set of devices that act as our own trusted personal assistant? 

## What it does
Imagine  having an Assistant that follows you, logging your activities as you go about your day.
Learning your habits, preferences and behaviour and thus "understands" you.

Assistant reminds you, or acts on your behalf, of important milestones,  
is a repository of your data,  
classified into personal and business areas,  
and keeps you safe by monitoring your environment.

It has you data securely encrypted residing on one or more local devices  
as well as in the cloud (available via a monthly subscription).

Your data is accessible only by you via multi-factor authentication with adaptive biometrics.
It even has a duress lockout if Assistant detects that you are being coerced against your will to reveal your data.

Assistant is available delivered as custom hardware
having cameras and microphones to constantly monitor your activities.
You don't even have to call to it to enable privacy mode -- it understands gestures.

It is also implemented as software agents installable on phones, tablets, computers, IP cameras and IP phones.
With an optional installation service being planned.

While Assistant is always a *personal* assistant,
Businesses can purchase a Business Add-on that allows lock out
of employee access to "Business" Assistant data.

There is a Wills / Power of attorney add-on planned to address the needs of the elderly.

For the Socially Active, there is a Dairy / Journal plan for users interested only in recording and searching personal activity logs.

For more info visit: https://aqimbo.beyondbroadcast.com

## How we built it
We used Go's multi-target capabilities in our Assistant custom device.
For power efficiency, this device will probably be based on ARM CPUs which Go fully supports.

There are also software agents written in Go for x86 computers,
the latest Apple ARM based computers, as well as ARM based tablets, phones and IP cameras.
These are also fully supported by Go.

Our software design was greatly simplified using Go's multitasking capabilities.
Each component / domain was implemented as autonomous goroutines.
These goroutines communicated with each other via a distributed messaging system (NATS/Jetstream).
NATS, being written in Go, was very easy to deploy as it is typically delivered as a single static binary.
However for aQimbo software, the NATS server is *embedded*, further simplifying deployment as there is only one binary to deploy.

Finally, our tiered Cloud service offerings run `Go` code to provide services.
Go's small binary footprint and efficient resource usage make it a smart
and efficient choice for cloud deployment.

We wish to acknowledge the contributions of Clare Zhuo for the artist impressions and rendering of the aQimbo portable custom hardware.

## Challenges we ran into
We thought hard about how to automate Business vs Personal data separation.
Our approach would probably use human-assisted, unsupervised learning classification.

We struggled with our portable device energy budget as it is always on.
This will probably be address with a accessory ecosystem (power bank stand, USB-C Power Delivery etc.).

The cameras on the custom device should have unhindered upward
as well as downward views. Downward views are important for fall-detection 
(elderly use case) and also to image-capture/OCR documents laying on a desk.
This requires the cameras to be offset from the base and thus make it taller and less portable.


## Accomplishments that we're proud of
We started with a well discussed set of requirements.

And implemented a proof-of-concept with simple Go code.

## What we learned
Building reliable distributed systems is a challenging task.
Components can and do get disconnected. Keeping data synchronized and updated can be difficult.

To address some of these challenges, we leveraged the capabilities of NATS/JetStream.
To make deployment simpler, we even embedded the NATS server into our code.
This resulted in a single, relatively small (15MB), static binary.

```
$ GOOS=linux GOARCH=arm64 go build -o /tmp/aqcustom cmd/aqimdemo/*.go

$ ls -lh /tmp/aqcustom 
-rwxr-xr-x 1 siuyin siuyin 15M May  8 01:42 /tmp/aqcustom

$ file /tmp/aqcustom 
/tmp/aqcustom: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked, not stripped
```

## What's next for Personal Assistant
aQimbo, the brand/company behind Personal Assistant is looking for startup funding.
