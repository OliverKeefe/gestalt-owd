package net.gestalt.app.service.token;

import io.smallrye.jwt.build.Jwt;
import io.smallrye.jwt.build.JwtClaimsBuilder;
import jakarta.enterprise.context.ApplicationScoped;
import net.gestalt.data.entity.Group;
import net.gestalt.data.entity.Organization;
import net.gestalt.data.entity.Role;
import net.gestalt.data.entity.User;

import java.time.Duration;
import java.util.stream.Collectors;

@ApplicationScoped
public class GenerateToken {

    protected String generate(User user) {
        return buildClaimsFromUser(user)
                .expiresIn(Duration.ofMinutes(15L))
                .sign();
    }

    private JwtClaimsBuilder buildClaimsFromUser(User user) {
        return Jwt.claims()
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
    }
}
