package net.gestalt.bootstrap.security.jwt;

import io.smallrye.jwt.build.Jwt;
import io.smallrye.jwt.build.JwtClaimsBuilder;
import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;
import net.gestalt.core.model.Group;
import net.gestalt.core.model.Organization;
import net.gestalt.core.model.Role;
import net.gestalt.core.model.User;

import java.util.stream.Collectors;

@ApplicationScoped
public class ClaimsBuilder {

    private final TokenProvider tokenProvider;

    @Inject
    public ClaimsBuilder(TokenProvider tokenProvider) {
        this.tokenProvider = tokenProvider;
    }

    public JwtClaimsBuilder fromUser(User user) {
        var claimsBuilder = Jwt.claims()
                .issuer("gestalt")
                .upn(user.getEmail())
                .subject(user.getEmail())
                .claim("userId", user.getId())
                .claim("role", user.getRole()
                        .stream()
                        .map(Role::getName)
                        .collect(Collectors.toSet())
                )
                .groups(user.getGroups()
                                .stream()
                                .map(Group::getName)
                                .collect(Collectors.toSet())
                )
                .claim("organizations", user.getOrganization()
                        .stream()
                        .map(Organization::getName)
                        .collect(Collectors.toSet())
                );

        return claimsBuilder;
    }
}
