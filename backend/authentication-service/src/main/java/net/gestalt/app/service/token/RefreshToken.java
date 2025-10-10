package net.gestalt.app.service.token;

import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;
import net.gestalt.data.entity.User;

@ApplicationScoped
public class RefreshToken {

    @Inject
    GenerateToken generateToken;

    @Inject
    VerifyToken verifyToken;

    private final User user;

    public RefreshToken(User user) {
        this.user = user;
    }

    protected String regenerate(String token) {
        try {
            verifyToken.isTokenValid(token);
            return generateToken.generate(user);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
