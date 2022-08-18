const { transform } = require("typescript");

//https://www.youtube.com/watch?v=3Z780EOzIQs

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
  theme: {
    extend: {
      colors: {
        blacky: "#000300",
        greeny: "#00df9a",
        darky: "#131315",
      },
      animation: {
        slide: "slide 36s linear infinite",
        customrotate: "customrotate",
      },
      keyframes: {
        slide: {
          "0%": { transform: "translate3d(0, 0, 0)" },
          "100%": { transform: "translate3d(-108rem, 0, 0)" },
        },
        customrotate: { transform: "rotate(180deg)" },
      },
    },
  },
  plugins: [],
};
