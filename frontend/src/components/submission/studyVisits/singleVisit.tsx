import { Question } from "@/services/clap/model/studyVisits";
import { Col, Container, Row } from "react-bootstrap";
import { useTranslation } from "react-i18next";
import SingleQuestion from "./singleQuestion";

function SingleVisit({ questions, submissionId }: { questions: Question[], submissionId: number }) {
  const { t } = useTranslation('submission/studyVisits');

  return (
    <Container>
      <Row className="font-bold border-1 border-lightgray">
        <Col>
          {t('question')}
        </Col>
        <Col>
          {t('answers')}
        </Col>
        <Col>
          {t('attachments')}
        </Col></Row>
      {questions.map((question, idx) => {
        return (
          <SingleQuestion question={question} submissionId={submissionId} key={`${submissionId}/${idx}`} className="border-1 border-lightgray p-2" />
        )
      })}
    </Container >
  )
}

export default SingleVisit;
