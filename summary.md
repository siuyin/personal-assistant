## Inspiration
Google assistant, Alexa and Siri are limited. My use of them is mostly as a glorified kitchen timer.

They are limited because they are sandboxed. Isolated because of privacy concerns.
Some of us keep journals or have lawyers act on our behalf.
What if we have a device or a set of devices that act as our trusted assistant? 

## What it does
Imagine you have an Assistant that follows you and logs your activities as you go about your day.
Learning your habits, preferences and behaviour and thus "understands" you.
Assistant reminds you (or acts on your behalf) of important milestones,
is a repository of your data classified into personal and business areas,
and keeps you safe by monitoring your environment.

Has you data securely encrypted residing only on the local device (initial purchase)
as well as in the cloud (monthly subscription).
Accessible only by you via multi-factor authentication (Duress lockout).

Assistant can be delivered as custom hardware
having cameras and microphones to constantly monitor your activities
(You can ask it to pause recording).

It is also implemented as software agents installable on phones, tablets, computers, IP cameras and IP phones.
An optional installation service will be offered.

While Assistant is always a *personal* assistant,
Businesses can purchase a Business Add-on that allows lock out
of employee access to "Business" Assistant data.

Wills / Power of attorney add-on (TODO)

Dairy / Journal plan for users interested only in recording and searching personal activity logs.

For more info visit: https://aqimbo.beyondbroadcast.com

## How we built it
We used Go's concurrency capabilities in the Assistant custom device.
For power efficiency this device is based on ARM CPUs which Go fully supports.

There are also software agents written in Go for x86 computers, again fully supported by Go.

Finally users on the Cloud service plan have Go code running in the cloud
making use of the rich capabilities available in the cloud environment.

## Challenges we ran into
Automated Business vs Personal data classification.

Power budgeting for portable devices.
Power supply source and connection as Assistant is always-on.

The cameras on the custom device should have unhindered upward
as well as downward views.
This requires the cameras to be offset from the base and thus make it less portable.
We addressed this with a low-cost charging/power supply stand to provide the offset.

## Accomplishments that we're proud of
We started with a well discussed set of requirements.

And implemented a proof-of-concept with simple Go code.

## What we learned

## What's next for Personal Assistant