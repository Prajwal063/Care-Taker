import express from 'express';
import {
    createService,
    deleteService,
    getAllServices,
    getServiceById,
    updateService,
} from '../controller/service.controller';
import { validate } from '../middleware';
import { CreateServiceType } from '../types';

const router = express.Router();

router
    .route('/')
    .get(getAllServices)
    .post(validate(CreateServiceType), createService);

router
    .route('/:id')
    .get(getServiceById)
    .put(validate(CreateServiceType), updateService)
    .delete(deleteService);

export default router;
