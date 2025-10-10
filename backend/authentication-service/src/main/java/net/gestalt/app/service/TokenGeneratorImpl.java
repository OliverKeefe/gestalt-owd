package net.gestalt.app.service;

import net.gestalt.adapter.persistence.repository.UserRepository;
import net.gestalt.core.port.TokenGeneratorPort;

public class TokenGeneratorImpl {
    private final TokenGeneratorPort tokenGenerator;
    private final UserRepository userRepository;

    public TokenGeneratorImpl(TokenGeneratorPort tokenGenerator, UserRepository userRepository) {
        this.tokenGenerator = tokenGenerator;
        this.userRepository = userRepository;
    }

    //public LoginResponse login()
}
