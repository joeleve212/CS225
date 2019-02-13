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

(cons 'a '(b c))
(cons 'a (cons 'b (cons 'c nil)))
(list 'a (cons 'b nil) (cons 'c nil))

(defun fourth (car (cdr (cdr (cdr '(1 2 3 4 5 6))))))
OUTPUT: 4 

(defun askem (string)
  (format t "~A" string)
  (read))
(if (> (x (read)) (y (read))) (x) (y))


;(defun greater (x y)
 ; (if (> x y) x y))
  ;(greater (2 1))

  
  
;QUESTION 10:

(write-line "Hello World")
v1:
(defun list-of-nums (n)
    (if (n==1) (1)
    ((cons nums n)
    (list-of-nums(n-1))))
    )
(let (x (read))
	(format t "~%" (list-of-nums x)))

	
v2:
(defun list-of-nums (n)
    (if (eql n 1) (1)
    ((cons nums n)
    (list-of-nums (- n 1))))
    )
(let (x (read))
	(format t "~%" (list-of-nums x)))
	
v3:
(defun nums (n)
    (format t "Running function")
    (if (eql n 1) (format t "~A" 1)
    (
        (cons n (numList))
        (nums (- n 1))))
    )
(format t "Test?~%")
(let (numList nil)
    (nums (read))
    )
