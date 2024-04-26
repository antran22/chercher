/** @type {import("tailwindcss").Config} */
module.exports = {
  content: ["./view/templates/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/forms"),
    require("@tailwindcss/typography"),
    require("daisyui"),
  ],

  daisyui: {
    themes: [
      {
        solarized: {
          primary: "#cb4b16",
          "primary-content": "#eee8d5",
          secondary: "#859900",
          accent: "#008aff",
          neutral: "#002b36",
          "base-100": "#fdf6e3",
          "base-200": "#eee8d5",
          "base-300": "#93a1a1",
          "base-content": "#002b36",
          info: "#268bd2",
          success: "#2aa198",
          warning: "#b58900",
          error: "#dc322f",
        },
      },
    ],
  },
};
