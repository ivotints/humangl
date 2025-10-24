# OpenGL Project — HumanGL

NOT FINISHED YET

Write `make run` to launch program

## Introduction

Since OpenGL 3.0 the native matrices and their associated functions (e.g. `glRotate`, `glPushMatrix`, etc.) are deprecated.

In this project you must implement your own matrix stack and matrix transformations in order to create a skeletal animation.

---

## Objectives

This project is an introduction to hierarchical modeling and matrix stack manipulation. You will learn to use matrices to link different parts of a model so they move together in a logical way.

---

## General instructions

* A **Makefile** or a similar build system is required.
* You may use the graphics library of your choice (SDL2, GLUT, SFML, etc.).
* You need to implement your own matrices and transformations and target **at least OpenGL 4.0**.
* You are free to use any programming language (I picked Go lang).

---

## Mandatory part

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
