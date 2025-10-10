import Keycloak from 'keycloak-js';

const keycloak = new Keycloak({
    url: 'http://127.0.0.1:8080',
    realm: 'gestalt',
    clientId: 'frontend-service',
});

let isInitialized = false;

export async function initKeycloak() {
    if (isInitialized) return keycloak;

    try {
        const authenticated = await keycloak.init({
            onLoad: 'login-required',
        });

        isInitialized = true;
        return authenticated ? keycloak : null;
    } catch (err) {
        console.error('Keycloak init failed:', err);
        return null;
    }
}

export default keycloak;