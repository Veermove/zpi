'use client'

import { useTranslation } from "@/app/i18n/client";
import { useAuthContext } from "@/context/authContext";
import { useClapAPI } from "@/context/clapApiContext";
import { PartialRating, RatingType } from "@/services/clap/model/rating";
import { redirect } from "next/navigation";
import { forwardRef, useEffect, useImperativeHandle, useState } from "react";
import { Button, Col, Container, Row } from "react-bootstrap";
import Error from "../error";
import Spinner from "../spinner";
import EditableRating from "./editableRating";

export type SingleRatingForwardData = {
  getIsReadyToSubmit: () => boolean;
}

interface SingleRatingProps {
  firstName: string;
  lastName: string;
  isEditable: boolean;
  criterionId?: number;
  ratingId?: number;
  isEditing: boolean;
  setIsEditing: (isEditing: boolean) => void;
  currentRating?: PartialRating;
  setCurrentRating: (currentRating: PartialRating) => void;
  assessorId: number;
  type: RatingType;
  draft: boolean;
}

const SingleRating = forwardRef<SingleRatingForwardData, SingleRatingProps>(({
  firstName,
  lastName,
  isEditable,
  criterionId,
  ratingId,
  isEditing,
  setIsEditing,
  currentRating,
  setCurrentRating,
  assessorId,
  type,
  draft
}, ref) => {
  const { t } = useTranslation('ratings/singleRating')
  const { user } = useAuthContext();
  const clapAPI = useClapAPI();
  const [error, setError] = useState('');

  useEffect(() => {
    if (!user) {
      redirect('/login');
    }
  });

  function buildErrorMessage(error: any): string {
    console.error(JSON.stringify(error))
    const message = error.response?.data?.errors?.[0];
    if (message && message.includes('length')) {
      return t('justificationLengthError');
    }
    return t('error');
  }

  const updateRating = async (justification: string, points: number) => {
    if (!clapAPI || !currentRating) {
      return;
    }

    try {
      // TODO current rating doesn't get updated
      setError('');
      const response = await clapAPI.upsertPartialRating({
        partialRatingId: currentRating.partialRatingId,
        justification,
        points
      });
      setCurrentRating({ ...currentRating, justification: response.justification, points: response.points });
      setIsEditing(false);
    } catch (error) {
      console.error('Unable to update rating:', error);
      setError(buildErrorMessage(error));
    }
  };

  const addRating = async (justification: string, points: number) => {
    setError('');
    if (!clapAPI || !assessorId || !criterionId || !ratingId) {
      return;
    }

    try {
      const response = await clapAPI.upsertPartialRating({
        justification,
        points,
        criterionId,
        ratingId,
      });
      setCurrentRating(response);
      setIsEditing(false);
    } catch (error) {
      console.error('Unable to add rating:', error);
      setError(buildErrorMessage(error));
    }
  }

  const onCancel = () => {
    setIsEditing(false);
    setError('');
  }

  useImperativeHandle(ref, () => ({
    getIsReadyToSubmit: () => {
      return !!currentRating && !!currentRating.justification;
    }
  }))

  if (!clapAPI) {
    return <Spinner />
  } else if (!isEditable && (!currentRating || draft)) { // no rating, current user has no permissions to add
    return <h6 className="text-purple my-4">{t('noRatingsFrom')} {firstName} {lastName}</h6 >
  } else {
    return (
      <Container className="text-left m-0 p-3">
        {error && (
          <Row>
            <Col>
              <Error text={error} />
            </Col>
          </Row>
        )}
        <Row>
          <Col xs={6} className="text-purple p-0">
            <h5>{t('assessor')} {firstName} {lastName}</h5>
          </Col>
          <Col xs={6} >
            {t(type)}
          </Col>
        </Row>
        {!currentRating && !!isEditing && ( // no rating, current user is adding it now
          <Col xs={12} >
            <EditableRating
              initialJustification=""
              initialPoints={0}
              onSubmit={addRating}
              onCancel={onCancel}
            />
          </Col>
        )}{!currentRating && !isEditing && ( // no rating, current user has permissions to add it
          <Row>
            <Col xs={10} />
            <Col xs={2}>
              <Button className="text-white" onClick={() => setIsEditing(true)}>{t('add')}</Button>
            </Col>
          </Row>
        )}{!!currentRating && !isEditing && ( // rating exists, in view mode
          <>
            <Row>
              <Col xs={2} className="text-purple font-bold">
                {currentRating.points} {t('points')}
              </Col>
              <Col xs={10}>
                {currentRating.justification}
              </Col>
            </Row>
          </>
        )}{!!currentRating && !!isEditable && !isEditing && ( // rating exists, in view mode, current user has permissions to edit
          <Row>
            <Col xs={10} />
            <Col xs={2} >
              <Button className="text-white" variant="primary" onClick={() => setIsEditing(true)}>
                {t('edit')}
              </Button>
            </Col>
          </Row>
        )}{!!currentRating && !!isEditable && !!isEditing && ( // rating exists, in edit mode, current user has permissions to edit
          <Row>
            <Col xs={12} >
              <EditableRating
                initialJustification={currentRating.justification}
                initialPoints={currentRating.points}
                onSubmit={updateRating}
                onCancel={onCancel}
              />
            </Col>
          </Row>
        )}
      </Container>
    )
  }
});

SingleRating.displayName = 'SingleRating';

export default SingleRating
