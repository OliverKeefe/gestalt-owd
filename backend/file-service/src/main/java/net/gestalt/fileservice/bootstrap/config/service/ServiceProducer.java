package net.gestalt.fileservice.bootstrap.config.service;

import jakarta.enterprise.context.ApplicationScoped;
import jakarta.enterprise.inject.Produces;
import net.gestalt.fileservice.core.service.BlobStorageServiceImpl;
import net.gestalt.fileservice.core.service.DistBlobStorageServiceImpl;
import net.gestalt.fileservice.core.service.IPFSHandlerImpl;
import net.gestalt.fileservice.core.service.MetadataServiceImpl;

@ApplicationScoped
public class ServiceProducer {

    @Produces
    @ApplicationScoped
    public MetadataServiceImpl metadataService() {
        return new MetadataServiceImpl();
    }

    @Produces
    @ApplicationScoped
    public BlobStorageServiceImpl blobStorageService() {
        return new BlobStorageServiceImpl();
    }

    @Produces
    @ApplicationScoped
    public DistBlobStorageServiceImpl distBlobStorageService() {
        return new DistBlobStorageServiceImpl();
    }

    @Produces
    @ApplicationScoped
    public IPFSHandlerImpl ipfsHandler() {
        return new IPFSHandlerImpl();
    }
}
