const { transform } = require("typescript");

//https://www.youtube.com/watch?v=3Z780EOzIQs

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      animation: {
        slide: "slide 36s linear infinite",
      },
      keyframes: {
        slide: {
          "0%": { transform: "translate3d(0, 0, 0)" },
          "100%": { transform: "translate3d(-108rem, 0, 0)" },
        },
      },
    },
  },
  plugins: [],
};
