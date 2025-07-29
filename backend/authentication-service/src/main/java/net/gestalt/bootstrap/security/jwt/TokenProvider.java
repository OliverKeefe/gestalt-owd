package net.gestalt.bootstrap.security.jwt;

import io.smallrye.jwt.build.Jwt;
import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;
import net.gestalt.core.model.User;
import net.gestalt.core.port.TokenGenerator;

import java.time.Duration;
import java.time.LocalDateTime;

@ApplicationScoped
public class TokenProvider implements TokenGenerator {

    @Inject
    ClaimsBuilder claimsBuilder;

    @Override
    public String generate(User user) {
        return claimsBuilder.fromUser(user)
                .expiresIn(Duration.ofMinutes(15L))
                .sign();
    }
}
