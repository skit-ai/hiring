# ML role for Programmers

> NOTE: The process is in development so things here might change.

> The idea of a pair programming call is inspired from a similar method used by
> [the Recurse Center](https://www.recurse.com/pairing-tasks).

[Job
description](https://angel.co/company/vernacular-ai/jobs/650173-machine-learning-role-for-software-engineers).
The interview process involves the following:

1. A phone screening round.
1. Programming round where we decide a problem beforehand and build it during a
   1 hour screen-shared call.

In this document, we keep a few notes for the programming round and a list of
tasks which can be proposed to the candidate for the call.

## Notes

* A `task` is defined as specification of a program, the basic version of which
  can be made within 30-40 mins. Basic version here doesn't have to perform
  efficiently or correctly for all cases. The code written just has to convey
  that the general direction and approach is correct if the project is taken
  ahead.
* Since this role is for Machine Learning, we try to cover a few ML problems
  holding each by their programming ruff so that no background in ML is required
  to finish the task. If any context is indeed needed, we make sure to provide
  that during or before the call. The idea is to have everything easy to
  conceptualize at a high level just by reading the description of each task
  without knowing ML.
* _Logistics_. We will do a screen sharing hangouts call which will be around 1
  hour. During/before the call, we will make sure to provide you with everything
  needed, both definitions/answers and starter code snippets/resources etc.
* _Completion_. The point is to see how the candidate plans the solution and
  goes ahead with the execution iteration by iteration. Most of which can be
  judged by being agnostic of whether the task was completed. Notice that
  _iteration_ is important as it helps us provide possibly valuable early
  feedback during the call and also to clarify expectations.
* _Scope_. We start with the scope defined in the task and keep expanding to
  fill in the time by building/thinking about other extensions. Dynamic scope
  expansion is an important piece of the process since that goes pretty far to
  tell how extensible or well designed the written program is.
* Tests, comments etc. While these are important for more persistent pieces of
  code, the candidate is free (and suggested) to skip these as we both are going
  to be have live mental models of the programs and can just poke around by
  questions/answers during the call.


## Task List

> The list is not final. Specifically, the following need work:
> 
>    1. prose quality
>    1. clarity in task description
>    1. more tasks

1. **Silence detection for audio streams**


A silence detection system is used by conversational systems to decide when to
stop listening to the user. A very basic system keeps a window of audio samples
and returns a boolean saying 'the user has stopped' once certain volume criteria
for last n seconds (say) is met. A simple criteria is to just threshold the
volume.

To elaborate on the samples bit, an audio stream is represented as samples each
of which can be something like 16 bit integer, 32 bit float etc. These values
are signed and represent volume (+ and - doesn't matter, the `abs` does).

The task is to make a program that takes a wav file with human speech and tells
at what time the user stopped speaking.

Not really needed, but [here](https://www.youtube.com/watch?v=FG9jemV1T7I) is a
nice introduction to digital media in general.

2. **Decoding words from pronunciations**

In speech recognition systems, a model (acoustic model) runs over the audio and,
effectively, returns a list of phonemes (fundamental units of sounds; equivalent
to characters for text). Each word is represented as a list of phonemes in what
is called a pronunciation dictionary. Here is an example:

```
// First token is word, next are the phoneme symbols showing
// how the word is said.
अँग्रेजी a mq g r ee j ii
अँधेरा a mq dh ee r aa
अँधेरे a mq dh ee r ee
अं a q
अंक a q k
अंकगणित a q k g a nx i t
अंकज्योतिषी a q k a j y o t i sx ii
अंकन a q k a n
अंकल a q k a l
अंकलेश्वर a q k l ee sh w a r

```

The task is to make a program which returns sequence of words given a sequence
of phonemes for a complete utterance (so no marked boundary between words) and a
dictionary mapping words to phonemes.

For example, for the above dictionary and the input `a mq g r ee j ii a q k a
l`, the output can be `अँग्रेजी अंकल`.

3. **Statistical Language Model**

At a high level, a language model (LM) tells the probability of a certain
sentence being more/less likely based on what the model has already seen.
Effectively the probability is low for texts which _surprise_ the model more.
For example consider the following sentences used for training:

```
i am a human
that is a cow
dogs are also human

```

If you use the LM trained on the texts above to find the probability of `cat` or
a phrase like `cat be hungry` you will get 0 since these were not even seen.

A simple way of making such models is to just do plain counting of tokens. These
are called statistical LMs. For a unigram LM, training essentially means
counting of all the single word tokens (thus uni-gram) in the whole training
corpus and getting probability by dividing by total tokens. For above sentences,
a unigram LM can be made using something like:

```python
counts = {
    "i": 1, "am": 1, "a": 2, "human": 2, "that": 1, "is": 1,
    "cow": 1, "dogs": 1, "are": 1, "also": 1
}
total_tokens = sum(counts.values())
lm = {k: v / total_tokens for k, v in counts.items()}

```

The task is to create a program which takes a corpus (maybe file with lines of
sentences) and creates an arbitrary n-gram LM (an n-gram LM counts not only
single tokens, but also phrases made up of upto n tokens).

4. **Composable rule parser**

Task here is to write a system that lets people create regex rules allowing
composition. For example, assuming a yaml representation, I want to be able to
write the following:

```
num:
  - \d
date:
  - <num> (Jan|Feb|...)
  - tomorrow|today
datetime:
  - on <date> at <num> (pm|am)

```

The program will take rules from such a file and provide a `parse` API which
might look like this:

```python
parse("hello world. let's meet on 3 Jan")
# [("date", "3 Jan")]

```


5. **Inferring conversational flows from cases**

For many conversational agent use cases, we create a _flow_ based on the problem
definition. A flow defines how conversations go in the sense that it tells what
does the bot do at every _state_ of the call. This can be seen as a finite state
machine (fsm) where we jump from state to state based on user's response.

Task here is to infer the structure of such a flow based on examples of real
conversations that happened using that flow. As an example, consider the
following two conversations:

```
BOT: hello
USER: get me a human
BOT: transferring

BOT: hello
USER: hello
BOT: bye

```

Looking at these two, I can infer the following basic flow:

```
hello
  + <u>get me a human
    + tranferring
  + <u>hello
    + bye

```

The program takes (two party) conversations in json and returns a possible json
representation of the underlying flow. We will use output from a simulator
[here](https://github.com/Vernacular-ai/ink-simulator) as conversations.

6. **Toy stitcher for text to speech**

Stitching in the sense of text to speech (TTS) means concatenating pieces of
audio segments to make full audio for a sentence.

For example, if we have a set of audio clips by a person saying digits from 0 to
9, we can make audio for a number like 35 (say) by just joining 3 and 5's audio.
The problem is that this approach is going to give 'three five' and not 'thirty
five'.

One way to solve this is to have a lot of random sequences recorded by the
person which hopefully cover most of the continuous sequences of such kinds.
This program works with such a pool of arbitrary sentences (represented as
phonemes; see task 2 for definitions) which supposedly have their audio files
somewhere accessible (we don't use these). Now when a new sentence comes in
(again, represented as a sequence of phonemes), the program tries to look up in
the pool, finds the best segments to stitch together and plays them
sequentially.

For our purpose, the program will only return a plan of stitching specifying
things like:

* pick `8` to `17` symbols from pool sentence 1
* then `3` to `5` from pool sentence 2
* and `1` to `1` from sentence 3.

A very simple approach will just keep on picking each phoneme separately, ask
whether that is available in the pool anywhere, append in the plan and move on.
A better program will see if it can find groups of phonemes already matching a
subsequence in the pool somewhere.
