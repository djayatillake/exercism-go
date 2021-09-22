
Uploaded avatar of andrerfcsantos
andrerfcsantos
20h ago
This is very good! Well done passing the tests.

This is the first time I've ever use concurrency in code!

You did an awesome job for your first time!

I also tried to use sync.Map but got horribly confused about how to iterate through the sync.Map to put the keys and values into a normal map afterwards. I understand that I'm supposed to use sync.Map.Range() for this but didn't understand how to use it. Is using sync.Map the right idea? It seemed to be from the articles suggested for reading in the Readme. It also then eliminates the need to reduce the list of maps afterwards which could scale badly.

sync.Map is a terrible idea here. There are two main reasons for this. The first is that if you are thinking on having a synchronized map of any type, this means you are also thinking about protecting each update to it with locks. This means that the program will spend a lot of time just trying to acquire / release locks without necessity. Your current approach where you have the goroutines computing submaps and then join them into a single one without using locks is a much more clean approach.

The second reason to not use sync.Map is that the implementation of it is like a map[interface{}]interface{}. This means that neither the keys or values have concrete types and can be anything! So, to work with this map, you'll have to be converting back and forth from your types to interface{} and vice-versa, which is just plain ugly. If you really want to have a synchronized map, the best option most of the time is to just create a regular map and protect it with a regular lock, like the example in Tour of Go: sync.Mutex.

Even the documentation of sync.Map warns about its usage:

The Map type is specialized. Most code should use a plain Go map instead, with separate locking or coordination, for better type safety and to make it easier to maintain other invariants along with the map content.

The documentation explains where sync.Map can be a better option than a regular map with a regular mutex and where it reduces lock contention. But they are very specific situations and one must think twice about choosing sync.Map.

I then made my second iteration after looking at some of the community solutions. Is this considered better than using the waitgroups? It does seem simpler.

I'd say that the use of a WaitGroup or not here should not be decided by its simplicity, but it depends more on the your perspective of the problem. Here you don't strictly need a WaitGroup, like this second iteration you made clearly shows. But one interesting exercise is to think why that is. Why don't you have to wait for all the goroutines here? And the reason it's because the number of sends and receives is well defined. Note that on line 19 you are iterating over s and in each iteration you create a goroutine that sends 1 time on the channel. And you know this loop will iterate len(s) times, so you know there will be len(s) sends on the channel. Then, when you are receiving you also iterate over s and in each iteration you receive 1 value, which means that in total you'll receive also len(s) values. Since you know the the number of sends and receives will be len(s) and your loops will iterate both len(s) times, you don't need any more special synchronization of goroutines.

Now, I want you to image a scenario where you don't know how many times your goroutines will send in the channel, but you still want to have some kind of loop performing the receive operations for you. One possiblity is to use a for range loop over a channel. Using the same variables you are using, this means that your second for loop could be something like:

  for m := range c {
    for k, v := range m {
      mf[k] += v
    }
  }
With this loop we are receiving maps m from the channel c, just like you are doing in your code. But notice how this loop is over the channel itself and not over s. But when will a loop over a channel terminate? Well, a loop over a channel ends when the channel is closed and there's nothing more to receive from it. Effectively, closing the channel is a signal we can send to the code receiving on the channel that nothing more will be sent on it. Ok, so now we must close the channel c for the for range loop over it terminate. But wait, when can we close the channel? Well, we can only close it when we know nothing more will be sent on it, and that will happen when we know all the goroutines sending on it are terminated. And when do we know all the goroutines are terminated? You guessed it - WaitGroup to the rescue!

What this means is that using a WaitGroup is kind of mandatory if you want to have a for range loop over a channel, because to have that you'll have to close the channel, and to close the channel you have to know when all the sending goroutines are done using a WaitGroup.

Of course, using WaitGroups is not a requirement on this exercise and your solution works very well. But as an extra challenge, I'd like to encourage you to try an implementation like the one I mentioned above, using a for range loop over a channel and using a WaitGroup to control when the channel is closed (hence, when that for range loop will terminate). That alternative solution might not be better performance-wise and even code-length wise, but I think you'll learn a thing or two implementing it.


Iteration 3was submitted
4h ago
Uploaded avatar of djayatillake
djayatillake
4h ago
I've had a go at what you have described, but I can't get it to work... it seems as though the goroutines never finish or at least never decrement the waitgroup counter.

It seems like I've followed quite closely some of the solutions I've seen online and it isn't depending on the channel closing.

https://gobyexample.com/waitgroups

I did also try using the defer method in the example but to the same result, I'm not sure why using the defer method would be better it seems more direct to explictly say waitgroup.Done() exactly where you want that to happen.

Uploaded avatar of andrerfcsantos
andrerfcsantos
3h ago
Hi!

This solution is a good step in the right direction. The problem you are having is exactly one of the reasons I wanted you to try this solution.

The tests fails because on line 23 because you are making a unbuffered channel. On these kinds of channels, a send will only happen when there's another goroutine receiving on it. The problem is that you only receive on the channel after your wg.Wait(). But for you to get pass that wait, you need for all the goroutines to finish. But for them to finish, you need to be able to send. And you can't send because you don't have any goroutine receiving it! This is known as a deadlock and no progress is made.

A quick-fix for this would be to make the channel be a buffered channel, with enough space for all the FreqMap we will be receiving:

c := make(chan FreqMap, len(s))
With this, all the goroutines will be able to send on the channel and finish, without the need to wait for another goroutine to be immediately ready to receive.

However, having a buffered channel with space for all the sends can be expensive, and we can make a smaller buffered channel, with space for 8 FreqMaps for instance:

c := make(chan FreqMap, 8)
With this, if we receive more than 8 strings, that will generate more than 8 FreqMaps to be sent on the channel, and we can get a deadlock as before. But we'll solve this shortly too, keep reading :)

Another thing that in this solution you are not doing, is closing the channel. Remember that after waiting for all the goroutines, we must close the channel:

wg.Wait()
close(c)
With all of this said, let's get back to the original problem: our real problem is that we want to close the channel after waiting for all the goroutines to be done, but at the same time we want to start receiving from the channel without necessarily waiting for all the goroutines to be done. So, what we really want is to move this bit:

wg.Wait()
close(c)
... out of the way of the main goroutine, so it can reach the code where it starts receiving from the channel, without waiting for all the other goroutines to finish. In Go, how do we move things "out of the way"?

Uploaded avatar of djayatillake
djayatillake
1h ago
I'm not sure I understand what you mean by "out of the way".

My initial thought when you said this was to put the wait and close in the loop over the channel but I didn't understand how that would help... I'm going to try putting it after the reduction loop over the channel and choosing an inadequate buffer, my thinking is that because the goroutines won't be able to finish the reduction loop will start taking a map out of the channel and using it, then allowing for a goroutine to continue etc until the end. Just tried it and it doesn't work

Uploaded avatar of andrerfcsantos
andrerfcsantos
1h ago
The trick here is to put the wait and closing of the channel in a separate goroutine, as in:

go func() {
  wg.Wait()
  close(c)
}()
With the wait logic on a separate goroutine, the main goroutine can now move on to receiving from the channel! This way we'll have 3 kinds of goroutines running concurrently:

goroutines computing FreqMaps and sending them on the channel
the main goroutine receiving from the channel, at the same the goroutines are sending them
a goroutine that's only waiting for all the sending goroutines to finish so it can close the channel. Closing the channel means the loop in the main goroutine over the channel that is receiving values can now finish.

Iteration 4was submitted
1h ago

Iteration 5was submitted
1h ago
Uploaded avatar of djayatillake
djayatillake
1h ago
https://dev.to/sophiedebenedetto/synchronizing-go-routines-with-channels-and-waitgroups-3ke2

I found the solution here, just as you were sending me that message.

It's interesting that you actually don't need to specify a buffer once the fix you described is implemented.

Is this worse though because it means each of the goroutines in the loop inputting to the channel is blocked until the channel is read from? Therefore, making the whole process serial rather than parallel?

Uploaded avatar of djayatillake
djayatillake
1h ago
with no buffer:

goos: darwin goarch: amd64 pkg: letter cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz BenchmarkConcurrentFrequency-8 49941 23854 ns/op 8242 B/op 52 allocs/op PASS ok letter 1.702s

with buffer of 3:

goos: darwin goarch: amd64 pkg: letter cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz BenchmarkConcurrentFrequency-8 40627 26049 ns/op 8268 B/op 53 allocs/op PASS ok letter 1.462s

Uploaded avatar of andrerfcsantos
andrerfcsantos
1h ago
Awesome!

It's interesting that you actually don't need to specify a buffer once the fix you described is implemented. Is this worse though because it means each of the goroutines in the loop inputting to the channel is blocked until the channel is read from? Therefore, making the whole process serial rather than parallel?

Yes exactly. Having a buffer, even a small one, can be enough to make sure all the goroutines send on the channel as fast as possible and can finish as fast as possible.

Take the benchmarks for this exercise as a pinch of salt, since the benchmarks only pass 3 strings to the function. The gains of such a strategy of having a buffered channel would show its benefits if we had more strings and hence more goroutines trying to send on the channel.

A last, small suggestion, would be to do wg.Add(len(s)) before the for loop instead of several wg.Add(1). WaitGroup is basically just a counter, that is protected by mutexes just so several goroutines can access it. Since each operation to the WaitGroup is protected by a mutex, each operation has to try to acquire a lock first and release the lock after the operation is finished, which adds some overhead compared to a +=1. So, ideally we want to perform as less Add operations as possible so we don't pay this overhead every time.


Iteration 6was submitted
1h ago
Uploaded avatar of djayatillake
djayatillake
44m ago
Thanks for your help! I've really feel like I understand this now.

This mentoring conversation is fit for stackoverflow!

Uploaded avatar of andrerfcsantos
andrerfcsantos
37m ago
Awesome!

Glad you took my little challenge of implementing an alternative solution despite already having one working. I feel like this was a very productive discussion indeed.

Thanks for being an awesome student!