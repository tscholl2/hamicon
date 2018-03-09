# hamicon

A small program to build randomized hamster icons.

They can be deterministcaly generated from a seed, so they can be used as user icons, or use random
ones as placeholders.

They are given in optimized-but-readable SVG format so they work for all sizes.
Without the CSS animations they are ~700bytes gziped.

They come with class names on the different body parts that are tied to css animations.
So you can have them wiggle their nose by adding the `.wiggle` class to the `#nose` component.

## Examples

![example 1](examples/example1.svg)

![example 2](examples/example2.svg)

![example 3](examples/example3.svg)

## Notes

To add CSS animations, send the `static/hamicon.css` file and add the corresponding classes.
There is some code to do this but it isn't setup quite yet.
A full example of all animations on a "blank" template is given in `static/example.svg`.
