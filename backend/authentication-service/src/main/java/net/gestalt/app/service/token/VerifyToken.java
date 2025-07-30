package net.gestalt.app.service.token;

import io.smallrye.jwt.auth.principal.ParseException;
import io.smallrye.jwt.build.Jwt;
import io.smallrye.jwt.auth.principal.JWTParser;
import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;
import org.eclipse.microprofile.jwt.Claims;
import org.eclipse.microprofile.jwt.JsonWebToken;

@ApplicationScoped
public class VerifyToken {

    @Inject
    JWTParser jwtParser;

    public boolean isTokenValid(String token) {
        try {
            JsonWebToken jwt = jwtParser.parse(token);
            return jwt.getIssuer().equals("gestalt") &&
                    !isExpired(jwt);
        } catch (ParseException e) {
            return false;
        }
    }

    private boolean isExpired(JsonWebToken jwt) {
        Long expires = jwt.getClaim(Claims.exp.name());
        return expires != null && expires * 1000 < System.currentTimeMillis();
    }

    private String getClaim(String token, String name) {
        try {
            JsonWebToken jwt = jwtParser.parse(token);
            return jwt.getClaim(name);
        } catch (ParseException e) {
            return null;
        }
    }

    private String getSubject(String token) {
        try {
            JsonWebToken jwt = jwtParser.parse(token);
            return jwt.getSubject();
        } catch (ParseException e) {
            return null;
        }
    }
}
