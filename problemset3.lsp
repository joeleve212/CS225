;;;Read Chapter 2 of Paul Graham ANSI Common Lisp. It is available online at:
;;;http://www.paulgraham.com/acl.html
;;;Scroll down past the book blurbs, chapters 1 and 2 are available near the bottom of the page.

;;;QUESTIONs 1 through 9
;;From the end of Chapter 2 do all the exercises. These are short exercises that you should first work on paper (or in your head), then you can check them with the Lisp interpreter.  gem125 or cslab recommended.

;;;Lists of numbers
;;;Question 10 - Write a lisp function (list­of­nums n) that will return a list of numbers (1 2 ... n )
;;;Question 11 Write a lisp function (remove­multiples k nlist) that will remove the multiples of k from a list of numbers nlist.
;;;so: (remove­multiples 3 (list­of­nums 6)) should produce (1 2 4 5).

;;;QUESTION 12
;;;Write a function (sieve n) that produces the prime numbers (and 1) up through n.  [Sound familiar?]


;;; Please submit your code as a text file, and a Word/libreoffice/etc document with screenshots of output,please.
;Work from earlier problems:
(cons 'a '(b c))
(cons 'a (cons 'b (cons 'c nil)))
(list 'a (cons 'b nil) (cons 'c nil))

(defun fourth (car (cdr (cdr (cdr '(1 2 3 4 5 6))))))
OUTPUT: 4 

(defun askem (string)
  (format t "~A" string)
  (read))
(if (> (x (read)) (y (read))) (x) (y))
  
;QUESTION 10:

(defun listofnums (x)
  (if (not (eql x 1)) 
    (listofnums (- x 1)) () )
  (format t "~A " x)
  )
(write-line "Input a number: ")
(listofnums (read))

;                              INCLUDE OUTPUT PICTURE FILE


;Question 11 version 1(works):

(defun removemultiples (k nlist)
    (format t "Inputs: ~A ~A ~%" k nlist)
    (remove k nlist)
    (loop for i in nlist do(
        if (eql 0 (mod i k));test
            (setf nlist (remove i nlist));then
            (format t "~A is not a multiple of ~A~%" i k);else
        )
    )
    (format t "~%New list: ~A" nlist)
)


;Version 2 (also works):
(defun removemultiples (k nlist)
    (loop for num in nlist
        do (
        if (eql (mod num k) 0)
            (setf nlist (remove num nlist))
            (format t "No change")
        )
        ; (return '(1 3))
        (print nlist)
    )
    (return nlist)
)

(format t "~A" (removemultiples 5 '(10 3 5 10 30 46)))

;Full Version of Q10 & Q11 together :

(defun removemultiples (k nlist)
    (format t "Inputs: ~A ~A ~%" k nlist)
    (loop for i in nlist do(
        if (eql 0 (mod i k));test
            (setf nlist (remove i nlist));then
            (format t "~A is not a multiple of ~A~%" i k);else
        )
    )
    (setf nlist nlist)
)

(defun listofnums (x)
    (setq numList nil)
    (loop
        (setq numList (cons x numList))
        (setf x (- x 1))
        (when (< x 1) (return numList))
    )
)

(format t "New List: ~A" (removemultiples 3 (listofnums 6)))

;       REMEMBER OUTPUT PIC FILE

; Question 12:

(defun sieve (n) 
    (setq primes nil)
    (setq x 1)

    (loop
        (setq flag 7); 7 = prime, 4 = non-prime (0 and 1 are boring)
        (loop for y from 2 to (/ x 2) do (
            if (= (mod x y) 0) 
                (setq flag 4)
            )
        )
        (if (= flag 7)
            (setq primes (cons x primes))
            (setq flag 7)
            )
        
        (setf x (+ x 1))
        (when (> x n) (return primes))
    )
)
 (format t "~A" (sieve 62))