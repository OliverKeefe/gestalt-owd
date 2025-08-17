package net.gestalt.fileservice.adapter.api.controller;

import jakarta.inject.Inject;
import jakarta.ws.rs.Consumes;
import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;
import net.gestalt.fileservice.adapter.api.dto.request.FileRequest;
import net.gestalt.fileservice.adapter.api.dto.response.FileResponse;
import net.gestalt.fileservice.bootstrap.mapper.FileMapper;
import net.gestalt.fileservice.core.model.FileModel;
import net.gestalt.fileservice.core.port.inbound.MetadataServicePort;

@Path("/files")
public class FileResource {

    @Inject
    MetadataServicePort metadataService;

    @Inject
    FileMapper fileMapper;

    @GET
    @Path("/get")
    @Consumes(MediaType.APPLICATION_JSON)
    @Produces(MediaType.APPLICATION_JSON)
    public Response getFile (FileRequest fileRequest) {
        FileModel fileModel;
        try {
            fileModel = fileMapper.toModel(fileRequest);
        } catch (Exception e) {
            // Log exception
            return Response.status(400, "Invalid Request").build();
        }

        FileResponse fileResponse = metadataService.getFile(fileModel);
        return Response.ok(fileResponse).build();
    }
}
