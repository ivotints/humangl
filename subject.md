# OpenGL Project — HumanGL

## Summary

This project is an introduction to hierarchical modeling.

---

## Chapter II — Introduction

Since OpenGL 3.0 the native matrices and their associated functions (e.g. `glRotate`, `glPushMatrix`, etc.) are deprecated.

In this project you must implement your own matrix stack and matrix transformations in order to create a skeletal animation.

---

## Chapter III — Objectives

This project is an introduction to hierarchical modeling and matrix stack manipulation. You will learn to use matrices to link different parts of a model so they move together in a logical way.

---

## Chapter IV — General instructions

* A **Makefile** or a similar build system is required. Only the contents of your repository will be evaluated.
* You may use the graphics library of your choice (SDL2, GLUT, SFML, etc.).
* You need to implement your own matrices and transformations and target **at least OpenGL 4.0**.
* You are free to use any programming language.

---

## Chapter V — Mandatory part

Body parts must be correctly articulated using your matrix stack. If the torso rotates, all the limbs must follow accordingly. If the upper arm moves, only the forearm should follow. When you modify the size of a limb, related parts must automatically reposition themselves.

Your model must contain the following parts (each drawn by a single function call — see constraints):

* a head
* a torso
* two arms, each with:

  * upper arm
  * forearm
* two legs, each with:

  * thigh
  * lower part

The model should be able to: walk, jump, and stay put.

---

## Chapter VI — Constraints

* Each body part must be drawn by **one and only one** function call. This function will draw a `1x1x1` geometric shape at the origin of the current matrix. Failure to follow this will cause loss of points.
* Upper and lower parts of the same limb are treated as distinct parts (i.e., separate draw calls).

---

## Chapter VIII — Submission and peer-evaluation

Turn in your assignment in your Git repository as usual. Only the work inside your repository will be evaluated during the defense. Double-check the names of your folders and files to ensure they are correct.

Be prepared to:

* Run the program and demonstrate the different movement patterns (walk, jump, idle).
* Modify limb sizes easily (either in code or at runtime).
* Show your drawing function, its calls, and explain how it works.
* Explain your hierarchical model and the resulting matrix stack.
