# Personal Assistant
`aQimbo` is a secure, learning, confidential personal assistant.

# Evaluating the demo app
A demonstration of this project is available as a ready to run docker image.
```
docker run --rm siuyin/aqimbodemo
```

then `docker ps` to find the running container. In the example below it is `nervous_elgamal`.
```
$ docker ps
CONTAINER ID   IMAGE               COMMAND                  CREATED         STATUS       PORTS  NAMES
2f0f5c005dc0   siuyin/aqimbodemo   "/nats-server -c /na…"   3 minutes ago   Up 3 minutes        nervous_elgamal
```

Obtain a bash shell in the running container thus
```
docker exec -it nervous_elgamal bash
```

Then run the demo app
```
$ /app
```
Your output should resemble:
```
INFO[0000] command centre started                        module=command
INFO[0000] learning system started                       module=learn
INFO[0000] location found: Singapore                     module=learn
INFO[0000] threat detected                               module=learn
INFO[0000] aQimbo custom hardware started                module=device
INFO[0000] phone bluetooth connected                     module=device
INFO[0000] save video for evidence requested             module=command
INFO[0000] police help requested                         module=command
```


# Requirements

## Feature: Personal Assistant
As a busy individual working in a modern business environment,  
I want a Personal Assistant that observes what I do and can act on my behalf  
so that I may focus on higher level work.

### Scenario: Keeps an activity log
Given I am a busy individual  
And I have Assistant monitoring my activities  
When I request for my activity log  
Then I should see a record of my activities.

### Scenario: Allows user to disable monitoring
Given I am an Assistant user  
When I do not want Assistant to monitor and log my activity  
Then I should see all Assistants pause their monitoring activities  
And I should get periodic reminders to re-enable monitoring.

### Scenario: Allows user to directly enter data
Given I am an Assistant user  
And I have an Assistant device with me  
When I wish to record and classify data  (example: who paid for lunch and the split between the attendees)  
Then I should see Assistant having captured that data.

### Scenario: Secures my data
Given I am an Assistant user  
And Assistant has collected confidential data  
When I test Assistant's security  
Then I should see Assistant's security upholds the most stringent security standards (example: multi-factor authentication, multi-stage authorisation as appropriate).

### Scenario: Duress lockout
Given I am an Assistant user
When Assistant detects I am under duress to instruct Assistant to reveal my personal data
Then I should see Assistant invoke a duress lockout procedure.

### Scenario: Able to categorise data as personal vs business data
Given I use Assistant in both a personal and a business (example: employee) context  
When I ask to separate personal and business data  
Then I should see a clear separation of data  
And when Assistant is not able to classify the data  
Then Assistant will ask me to clarify on a periodic basis as these ambiguities arise.

### Scenario: Business Add-on controls employee access to Business Data
Given I use Assistant in both a personal and a business (example: employee) context  
And my employer has purchased the Business Add-on
When I leave that employer
Then I should no longer have access to Business data on my Assistant.

### Scenario: Business Add-on has access to all Business classified data
Given I use Assistant in both a personal and a business (example: employee) context  
And I have installed Assistant software on my computing device used for business
When my employer wants to evaluate my interaction with my computing device and hence the rest of the business
Then my employer can get a summary of my activities.

### Scenario: Responds to an attention word, phrase, action or gesture
Given I am an Assistant user  
And I want Assistant to pay special attention (example: to record an important fact)  
When I speak a custom word or phase (eg. aqimbo)  or make a predefined gesture or push a designated button  
Then I should see Assistant give additional emphasis to the subsequent events.

### Scenario: Learns my activities
Given I am a Assistant user  
And I have an Assistant that has recorded my activity log  
When I ask for recommendations on a given topic  
Then I should see a list of recommendations listed in relevance order based on the patterns in the data of my activity log (example: driven by unsupervised learning / classification).

### Scenario: Acts to safeguard my well-being
Given I am an Assistant user  
When an Assistant detects a threat to my safety  
Then I should see that Assistant has called for assistance.

### Scenario: Act on my commands
Given I am a busy individual  
And I have an Assistant watching / listening for my commands  
When I ask Assistant to perform an action  
Then I should see my action being carried out  
But if Agent is unable to carry out my command (example: pay for a purchase with my credit card, or query the balance in my bank account)  
Then I should be notified  
And Agent should recommend alternative options where available.

### Scenario: Helps me manage my deadlines
Given I am an Assistant user  
And Assistant knows of my significant milestones and deadlines 
When a milestone or deadline approaches  
Then I should be reminded by Assistant  
Or if I have previously authorised an action  
Then I should see Assistant carry out that action (example: auto-renew Assistant service)  .

### Scenario: Helps me with automated reporting.
Given I am an Assistant user  
When I call for a summary of my activities over a specified period  
Then I should see a report with automatically classified data and their summary statistics.

### Scenario: Assistant is available in multiple places
Given I am an Assistant user  
When I call for my Assistant in multiple locations  
Then I should receive an acknowledgement from my Assistant.

### Scenario: Assistant learns it location
Given I am an Assistant user
And I have an Assistant device with me
When I check Assistant's recognised locations
Then I should see a list of locations, some which will have been labelled by User through interaction with Assistant.

### Scenario: Assistant is available on a custom portable device
Given I am an Assistant user  
When I need Assistant's services  
Then I should be able to carry an Assistant device with me  
And that device should contain a rechargeable power source which can be charged with readily available power supplies.

### Scenario: Able to capture a video (and audio) log
Given I am an Assistant user  
And I have my Assistant custom portable device with me  
Or I have Assistant software agent installed on a video camera (example IP camera system)  
When I ask Assistant for a visual summary of my activities  
Then I should see that summary presented on a display.

### Scenario: Able to capture a interactions with computing and communication devices
Given I am an Assistant user  
And Assistant software agent is installed on the computing (example: computer)  or communication (example IP phone) device I interact with  
When I ask Assistant for a log of my interactions with those devices  
Then I should see that log presented on a display.

### Scenario: Assistant is available on multiple devices
Given I am an installation technician  
When I am assigned to install the Assistant software agent  
Then I should be able the Assistant on multiple devices.

### Scenario: Assistant is workable with an intermittent internet connection
Given I am an Assistant user  
And I interact with Assistant when Assistant has no internet connectivity  
When I bring Assistant to an environment with internet connectivity  
Then I should see Assistant upload and synchronize data across all my Assistants and carry out actions that require an internet connection.

### Scenario: Assistant is able to report its own health status
Given I am an Assistant user  
When I interact with Assistant  
Then I should be alerted by Assistant when it needs attention (example: low battery or out if internet range).

# Event Storming / Context Mapping
[whiteboard link](https://miro.com/app/board/o9J_ldKJFc4=/?share_link_id=330109991457).

# Building the demo app
To build the docker image locally, do the following:

Clone this repository:
```
git clone https://github.com/siuyin/personal-assistant
```

Build the image:
```
docker build -t aqimbodemo
```

Run the image:
```
docker run -it --rm aqimbodemo
```
