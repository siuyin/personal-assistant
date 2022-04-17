# Personal Assistant

## Requirements

### Feature: Personal Assistant
As a busy individual working in a modern business environment,
I want a Personal Assistant that observes what I do and can act on my behalf
so that I may focus on higher level work.

#### Scenario: Keep an activity log
Given I am a busy individual
And I have a software agent monitoring my activities
When I request for my activity log
Then I should see a record of my activities. 

#### Scenario: Act on my commands
Given I am a busy individual
And I have a software Agent watching / listening for my commands
When I ask Agent to perform an action
Then I should see my action being carried out
But if Agent is unable to carry out my command
Then I should be notified
And Agent should recommend alternative options if available.


