package net.gestalt.app.service.token;

import jakarta.inject.Inject;
import jakarta.inject.Singleton;
import net.gestalt.data.entity.User;

@Singleton
public class TokenService {

    @Inject
    GenerateToken generateToken;

    public String generateToken(User user) {
        try {
            return generateToken.generate(user);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

}
