package net.gestalt.bootstrap.security.jwt;

import io.smallrye.jwt.build.Jwt;
import io.smallrye.jwt.build.JwtClaimsBuilder;
import jakarta.enterprise.context.ApplicationScoped;
import net.gestalt.core.model.User;

@ApplicationScoped
public class ClaimsBuilder {

    public JwtClaimsBuilder fromUser(User user) {
        var claimsBuilder = Jwt.claims()
                .issuer("gestalt")
                .upn(user.getEmail())
                .subject(user.getEmail())
                .claim("userId", user.getId())
                .claim("role", user.getRole())
                .groups("groups", user.getGroups().getName());

        if (user.getOrganization() != null) {
            claimsBuilder.claim("organizationId", user.getOrganization().getId());
        }

        return claimsBuilder;
    }
}
