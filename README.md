# Personal Assistant

## Requirements

### Feature: Personal Assistant
As a busy individual working in a modern business environment,  
I want a Personal Assistant that observes what I do and can act on my behalf  
so that I may focus on higher level work.

#### Scenario: Keeps an activity log
Given I am a busy individual  
And I have Assistant monitoring my activities  
When I request for my activity log  
Then I should see a record of my activities.

#### Scenario: Allows user to disable monitoring
Given I am an Assistant user  
When I do not want Assistant to monitor and log my activity  
Then I should see all Assistants pause their monitoring activities  
And I should get periodic reminders to re-enable monitoring.

#### Scenario: Allows user to directly enter data
Given I am an Assistant user  
And I have an Assistant device with me  
When I wish to record and classify data  (example: who paid for lunch and the split between the attendees)  
Then I should see Assistant having captured that data.

#### Scenario: Secures my data
Given I am an Assistant user  
And Assistant has collected confidential data  
When I test Assistant's security  
Then I should see Assistant's security upholds the most stringent security standards (example: multi-factor authentication, multi-stage authorisation as appropriate).

#### Scenario: Duresss lockout
Given I am an Assistant user
When Assitant detects I am under duress to instruct Assistant to reveal my personal data
Then I should see Assistant invode a duress lockout procedure.

#### Scenario: Able to categorise data as personal vs business data
Given I use Assistant in both a personal and a business (example: employee) context  
When I ask to separate personal and business data  
Then I should see a clear separation of data  
And when Assistant is not able to classify the data  
Then Assistant will ask me to clarify on a periodic basis as these ambiguities arise.

#### Scenario: Responds to an attention word, phrase, action or gesture
Given I am an Assistant user  
And I want Assistant to pay special attention (example: to record an important fact)  
When I speak a custom word or phase (eg. greta-bee)  or make a predefined gesture or push a designated button  
Then I should see Assistant give additional emphasis to the subsequent events.

#### Scenario: Learns my activities
Given I am a Assistant user  
And I have an Assistant that has recorded my activity log  
When I ask for recommendations on a given topic  
Then I should see a list of recommendations listed in relevance order based on the patterns in the data of my activity log.

#### Scenario: Acts to safeguard my well-being
Given I am an Assistant user  
When an Assistant detects a threat to my safety  
Then I should see that Assistant has called for assistance.

#### Scenario: Act on my commands
Given I am a busy individual  
And I have an Assistant watching / listening for my commands  
When I ask Assistant to perform an action  
Then I should see my action being carried out  
But if Agent is unable to carry out my command (example: pay for a purchase with my credit card)  
Then I should be notified  
And Agent should recommend alternative options where available.

#### Scenario: Helps me manage my deadlines
Given I am an Assistant user  
And Assistant knows of my significant milestones and deadlines 
When a milestone or deadline approaches  
Then I should be reminded by Assistant  
Or if I have previously authorised an action  
Then I should see Assistant carry out that action (example: auto-renew Assistant service)  .

#### Scenario: Assistant is available in multiple places
Given I am an Assistant user  
When I call for my Assistant in multiple locations  
Then I should receive an acknowledgement from my Assistant.

#### Scenario: Assistant is available on a custom portable device
Given I am an Assistant user  
When I need Assistant's services  
Then I should be able to carry an Assistant device with me  
And that device should contain a rechargeable power source which can be charged with readily available power supplies.

#### Scenario: Able to capture a video (and audio) log
Given I am an Assistant user  
And I have my Assistant custom portable device with me  
Or I have Assistant software agent installed on a video camera (example IP camera system)  
When I ask Assistant for a visual summary of my activities  
Then I should see that summary presented on a display.

#### Scenario: Able to capture a interactions with computing and communication devices
Given I am an Assistant user  
And Assistant software agent is installed on the computing (example: computer)  or communication (example IP phone) device I interact with  
When I ask Assistant for a log of my interactions with those devices  
Then I should see that log presented on a display.

#### Scenario: Assistant is available on multiple devices
Given I am an installation technician  
When I am assigned to install the Assistant software agent  
Then I should be able the Assistant on multiple devices.

#### Scenario: Assistant is workable with an intermittent internet connection
Given I am an Assistant user  
And I interact with Assistant when Assistant has no internet connectivity  
When I bring Assistant to an environment with internet connectivity  
Then I should see Assistant upload and synchronize data across all my Assistants and carry out actions that require an internet connection.

#### Scenario: Assistant is able to report its own health status
Given I am an Assistant user  
When I interact with Assistant  
Then I should be alerted by Assistant when it needs attention (example: low battery or out if internet range).
