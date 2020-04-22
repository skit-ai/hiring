
# Platform Frontend Intern Role

We at Vernacular.ai use various tools to support our CCA and ASR platform, some of which are interfaces that we have created for ourselves designed specifically for our use-case.

### Problem Statement
Your task is to create an interface which consists of the following:
 
-   A card that has two input boxes where a user can enter text.
-   The user will then be able to record the statement he has given as inputs.
-   The interface will provide the user with an option to listen to what they have recorded and re-record it.
-   A card here can be viewed as a snippet of a conversation between two people.
-   The interface will allow users to add multiple snippets by providing a button on the screen to add more snippets.
-   The interface will also give an option to delete a snippet.
-   Input box 1 and its recording will be considered to be the “BOT”, input box 2 and its recording will be considered to be a “Customer”.
-   The interface now allows these audios recorded in the snippets to be stitched together so that it appears as if the Bot and the customer are speaking on a call.
-   The interface will have an option to preview the audios stitched together without downloading the stitched audio file.
-   The interface will allow the user to download the generated audio file, along with the transcription which is a JSON file, You can decide what all metadata you want to save in this file apart from the text.
-   For the stitched audio the interface will add 1500ms of silence between the Customer and the Bot and will add 750ms between Bot and Customer.

#### Bonus Point
- Snippets can be swapped. 
- Styling. Using any CSS framework is highly discouraged.
- Hosting it on any service.

### Instructions for the Candidate
- Please note we expect this challenges to be completed with either [ReactJS](https://reactjs.org/) or [Elm Lang](https://elm-lang.org/). Angular, Vue.js and Ember or even VanillaJs will not be entertained.
- Please ensure your code is available on Github.
- Documentation explaining how to use the Interface.

### Evaluation Criteria
- Completeness of the question
- Code quality and use of a linter
- Git history

### Reading Links
- https://guide.elm-lang.org/
