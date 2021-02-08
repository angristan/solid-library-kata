# Solid Library Kata

We decided to try an experiment in a new Field: software for Libraries.
In order to build this proof of concept, we interviewed a few potential clients.
This is the summary of the most important features to start with:

* For the books we have in our library, we identify them by their title and by
  the author's name.
* Only the librarian can add and reference a new book in the library
* Guests can see the content of the library but they cannot borrow books.
* Members can borrow books but they can only have 3 books at the same time.
  They can only keep the books for a maximum period of time of 4 weeks _(ex: They can
  borrow 2 books, return 1 after 1 week, take 2 another books, return the first
  after 3 more weeks and the 2 new books after 4 weeks)_
* a user only needs an unique login to be identified (you don't need names, passwords
  and so on for this version)



Technical notes:

- There is no bi-directional user interaction but your app needs to be scriptable.
- you can choose the persistence of your choice, but the simpler the better
- your application needs to be testable and must be tested

Design notes:

- the application should be very simple in terms of user interface (ie: an
  app console is good enough for this test) but internally well designed.
- ask questions to check if there is hidden rules
- define your main use cases and find 2-3 exemples to illust