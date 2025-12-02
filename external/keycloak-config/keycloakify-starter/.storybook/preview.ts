import type { Preview } from "@storybook/react-vite";

const preview: Preview = {
    parameters: {
        controls: {
            matchers: {
                color: /(background|color)$/i,
                date: /Date$/i
            }
        }
    }
};

// .storybook/preview.js

// REMOVE PatternFly
document.addEventListener("DOMContentLoaded", () => {
    [...document.querySelectorAll("link")].forEach(link => {
        if (link.href.includes("patternfly") || link.href.includes("keycloak")
        ) {
            link.remove();
        }
    });
});


export default preview;
